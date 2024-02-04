package order_infrastructure

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (cc OrderRepository) GetTotalAmountSpent(orderHistories []primitive.ObjectID) (float64, error) {
	// 集計パイプラインの作成
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: orderHistories}}}}}},
		{{Key: "$group", Value: bson.D{{Key: "_id", Value: nil}, {Key: "totalAmount", Value: bson.D{{Key: "$sum", Value: "$total_amount"}}}}}},
	}

	// 集計パイプラインの実行
	cursor, err := cc.Collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return 0, errors.Wrap(err, "cc.Collection.Aggregate")
	}
	defer cursor.Close(context.TODO())

	var result struct {
		TotalAmount float64 `bson:"totalAmount"`
	}

	if cursor.Next(context.TODO()) {
		if err := cursor.Decode(&result); err != nil {
			return 0, errors.Wrap(err, "cursor.Decode")
		}
	} else {
		// カーソルにデータがない場合、合計金額は0とする
		return 0, nil
	}

	return result.TotalAmount, nil
}
