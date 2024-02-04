package order_usecase

import (
	"go-mongodb-sample/app/infrastructures"
	customer_infrastructure "go-mongodb-sample/app/infrastructures/customers"
	order_infrastructure "go-mongodb-sample/app/infrastructures/orders"
	model "go-mongodb-sample/app/models"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (o OrderService) Create(tm *infrastructures.MongoTransactionManager, cc *customer_infrastructure.OrderRepository, dto CreateDTO) error {
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
		return errors.Wrap(err, "cc.FindByID")
	}

	// トランザクションを使用するためのセッションを開始
	session, err := tm.Client.StartSession()
	if err != nil {
		return errors.Wrap(err, "client.StartSession")
	}
	defer session.EndSession(o.Ctx)

	// トランザクションを開始
	if err = mongo.WithSession(o.Ctx, session, func(sc mongo.SessionContext) error {
		cc.Ctx = sc
		oi := order_infrastructure.NewOrderRepository(sc, tm.Client.Database(o.DBName))

		if err := sc.StartTransaction(); err != nil {
			return errors.Wrap(err, "session.StartTransaction")
		}
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
		createdOrder, err := oi.Create(newOrder)
		if err != nil {
			return errors.Wrap(err, "oi.Create")
		}

		// カスタマーのオーダー履歴を追加する
		if err = cc.UpdateHistory(dto.CustomerID, createdOrder.ID); err != nil {
			return errors.Wrap(err, "cc.UpdateHistory")
		}

		// コミット
		if err = sc.CommitTransaction(sc); err != nil {
			return errors.Wrap(err, "session.CommitTransaction")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "mongo.WithSession")
	}

	return nil
}
