package order_infrastructure

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (cc OrderRepository) GetTotalAmountSpent(orderHistories []bson.ObjectID) (float64, error) {
	// 集計パイプラインの作成
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: orderHistories}}}}}},
		{{Key: "$group", Value: bson.D{{Key: "_id", Value: nil}, {Key: "totalAmount", Value: bson.D{{Key: "$sum", Value: "$total_amount"}}}}}},
	}

	// 集計パイプラインの実行
	cursor, err := cc.Collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return 0, fmt.Errorf("cc.Collection.Aggregate: %w", err)
	}
	defer cursor.Close(context.TODO())

	var result struct {
		TotalAmount float64 `bson:"totalAmount"`
	}

	if cursor.Next(context.TODO()) {
		if err := cursor.Decode(&result); err != nil {
			return 0, fmt.Errorf("cursor.Decode: %w", err)
		}
	} else {
		// カーソルにデータがない場合、合計金額は0とする
		return 0, nil
	}

	return result.TotalAmount, nil
}
