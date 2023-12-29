package order_usecase

import (
	"context"
	customer_infrastructure "go-mongodb-sample/app/infrastructures/customers"
	order_infrastructure "go-mongodb-sample/app/infrastructures/orders"
	model "go-mongodb-sample/app/models"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (o OrderService) Create(dto CreateDTO) error {
	oi := order_infrastructure.NewOrderRepository(o.Ctx, o.DB.Collection("orders"))
	cc := customer_infrastructure.NewCustomerRepository(o.Ctx, o.DB.Collection("customers"))

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

	// FIXME: ユースケース層がmongoDBのドライバーに依存している
	// clientを作成
	client, err := mongo.Connect(o.Ctx, options.Client().ApplyURI(o.ConnectionString))
	if err != nil {
		return errors.Wrap(err, "mongo.Connect")
	}
	defer client.Disconnect(o.Ctx)

	// トランザクションを使用するためのセッションを開始
	session, err := client.StartSession()
	if err != nil {
		return errors.Wrap(err, "client.StartSession")
	}
	defer session.EndSession(context.Background())

	// トランザクションを開始
	if err = mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {
		if err := session.StartTransaction(); err != nil {
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

		return nil
	}); err != nil {
		return errors.Wrap(err, "mongo.WithSession")
	}

	return nil
}
