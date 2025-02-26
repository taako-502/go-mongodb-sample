package order_usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure/customer_infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/order_infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/transaction_manager"
	"github.com/taako-502/go-mongodb-sample/app/model"
)

func (o orderService) Create(tm transaction_manager.TransactionManager, co model.CustomerAdapter, oi model.OrderAdapter, dto CreateDTO) error {

	// dtoからmodelを作成する
	detailsModel := make([]model.OrderDetail, len(dto.OrderDetails))
	for i, v := range dto.OrderDetails {
		detailsModel[i] = *model.NewOrderDetail(v.ProductID, v.Quantity, v.Price)
	}
	model, err := model.NewOrder(dto.CustomerID, detailsModel, dto.OrderDate, dto.Status)
	if err != nil {
		return fmt.Errorf("model.NewOrder: %w", err)
	}

	// カスタマーが存在するか確認する
	if _, err := co.FindOne(dto.CustomerID); err != nil {
		if errors.Is(err, customer_infrastructure.ErrCustomerNotFound) {
			return ErrCustomerNotFound
		}
		return fmt.Errorf("customer_infrastructure.OrderRepository.FindByID: %w", err)
	}

	// トランザクションを開始
	if err = tm.WithSession(func(ctx context.Context) error {
		// NOTE: トランザクション内のテストができていない
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
			return fmt.Errorf("order_infrastructure.OrderRepository.Create: %w", err)
		}

		// カスタマーのオーダー履歴を追加する
		if err = co.UpdateHistory(dto.CustomerID, createdOrder.ID); err != nil {
			return fmt.Errorf("customer_infrastructure.OrderRepository.UpdateHistory: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("tm.WithSession: %w", err)
	}

	return nil
}
