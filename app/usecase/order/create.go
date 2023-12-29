package order_usecase

import (
	"context"
	customer_infrastructure "go-mongodb-sample/app/infrastructures/customers"
	order_infrastructure "go-mongodb-sample/app/infrastructures/orders"
	model "go-mongodb-sample/app/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderDetailDTO struct {
	ProductID primitive.ObjectID
	Quantity  int
	Price     float64
}

type CreateDTO struct {
	CustomerID   primitive.ObjectID
	OrderDetails []OrderDetailDTO
	OrderDate    time.Time
	Status       string
}

func Create(ctx context.Context, DB *mongo.Database, dto CreateDTO) error {
	// FIXME: ここでcollectionのインスタンスを作成するとユースケース層がDBに依存してしまう
	oi := order_infrastructure.NewOrderRepository(ctx, DB.Collection("orders"))
	cc := customer_infrastructure.NewCustomerRepository(ctx, DB.Collection("customers"))
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
	// NOTE: ここでトランザクションがあるとよい
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
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// カスタマーの履歴を更新する
	err = cc.UpdateHistory(dto.CustomerID, createdOrder.ID)
	if err != nil {
		return errors.Wrap(err, "cc.UpdateHistory")
	}
	return nil
}
