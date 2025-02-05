package product_controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/product_infrastructure"
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/labstack/echo/v4"
)

func (pc ProductController) FindOne(c echo.Context) error {
	ID, err := bson.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// コンテキストを設定
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbm, err := infrastructure.NewMongoDBManager(ctx, pc.ConnectionString)
	if err != nil {
		return fmt.Errorf("NewMongoDBManager: %w", err)
	}
	defer dbm.Client.Disconnect(ctx)

	pi := product_infrastructure.NewProductRepository(ctx, dbm.Client.Database(pc.DBName))
	order, err := pi.FindOne(ID)
	if errors.Is(err, product_infrastructure.ErrProductNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}
