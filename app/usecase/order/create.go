package order_usecase

import (
	"context"
	"go-mongodb-sample/app/infrastructure"
	customer_infrastructure "go-mongodb-sample/app/infrastructure/customers"
	order_infrastructure "go-mongodb-sample/app/infrastructure/orders"
	"go-mongodb-sample/app/model"

	"github.com/pkg/errors"
)

func (o OrderService) Create(tm *infrastructure.MongoTransactionManager, cc *customer_infrastructure.OrderRepository, dto CreateDTO) error {
	// dtoからmodelを作成する
	detailsModel := make([]model.OrderDetail, len(dto.OrderDetails))
	for i, v := range dto.OrderDetails {
		detailsModel[i] = *model.NewOrderDetail(v.ProductID, v.Quantity, v.Price)
	}
	model, err := model.NewOrder(dto.CustomerID, detailsModel, dto.OrderDate, dto.Status)
	if err != nil {
		return errors.Wrap(err, "model.NewOrder")
	}

	// カスタマーが存在するか確認する
	if _, err := cc.Find(dto.CustomerID); err != nil {
		if errors.Is(err, customer_infrastructure.ErrCustomerNotFound) {
			return ErrCustomerNotFound
		}
		return errors.Wrap(err, "customer_infrastructure.OrderRepository.FindByID")
	}

	// トランザクションを使用するためのセッションを開始
	session, err := tm.StartSession()
	if err != nil {
		return errors.Wrap(err, "tm.StartSession")
	}
	defer session.EndSession(o.Ctx)

	// トランザクションを開始
	if err = tm.WithSession(o.Ctx, session, func(ctx context.Context) error {
		// TODO: 在庫数を更新する
		// オーダーを永続化する
		var totalAmount float64
		newOrderDetails := make([]order_infrastructure.OrderDetailDTO, len(model.OrderDetails))
		for i, v := range model.OrderDetails {
			d := order_infrastructure.NewOrderDetailDTO(v.ProductID, v.Quantity, v.Price)
			newOrderDetails[i] = *d
			totalAmount += v.Price * float64(v.Quantity)
		}
		newOrder := order_infrastructure.NewOrderDTO(
			model.CustomerID,
			newOrderDetails,
			model.OrderDate,
			totalAmount,
			model.Status,
		)
		oi := order_infrastructure.NewOrderRepository(ctx, tm.Client.Database(o.DBName))
		createdOrder, err := oi.Create(newOrder)
		if err != nil {
			return errors.Wrap(err, "order_infrastructure.OrderRepository.Create")
		}

		// カスタマーのオーダー履歴を追加する
		if err = cc.UpdateHistory(dto.CustomerID, createdOrder.ID); err != nil {
			return errors.Wrap(err, "customer_infrastructure.OrderRepository.UpdateHistory")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "tm.WithSession")
	}

	return nil
}
