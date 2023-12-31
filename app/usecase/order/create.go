package order_usecase

import (
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
	// dtoからmodelを作成する
	detailsModel := make([]model.OrderDetail, len(dto.OrderDetails))
	for i, v := range dto.OrderDetails {
		detailsModel[i] = *model.NewOrderDetail(v.ProductID, v.Quantity, v.Price)
	}
	model, err := model.NewOrder(dto.CustomerID, detailsModel, dto.OrderDate, dto.Status)
	if err != nil {
		return errors.Wrap(err, "model.NewOrder")
	}

	// FIXME: ユースケース層がmongoDBのドライバーに依存している
	// clientを作成
	client, err := mongo.Connect(o.Ctx, options.Client().ApplyURI(o.ConnectionString))
	if err != nil {
		return errors.Wrap(err, "mongo.Connect")
	}
	defer client.Disconnect(o.Ctx)

	// カスタマーが存在するか確認する
	cc := customer_infrastructure.NewCustomerRepository(o.Ctx, client.Database(o.DBName))
	if _, err := cc.Find(dto.CustomerID); err != nil {
		return errors.Wrap(err, "cc.FindByID")
	}

	// トランザクションを使用するためのセッションを開始
	session, err := client.StartSession()
	if err != nil {
		return errors.Wrap(err, "client.StartSession")
	}
	defer session.EndSession(o.Ctx)

	// トランザクションを開始
	if err = mongo.WithSession(o.Ctx, session, func(sc mongo.SessionContext) error {
		cc.Ctx = sc
		oi := order_infrastructure.NewOrderRepository(sc, client.Database(o.DBName).Collection("orders"))

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
