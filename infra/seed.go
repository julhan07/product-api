package infra

import (
	"api-product/entities"
	"context"
	"math/rand"

	"github.com/bxcodec/faker/v4"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

var ctx = context.Background()

func RunSeed(dbSource, dbDesti *gorm.DB, count int, redisConn *redis.Client) error {

	var sourceProduct entities.ProductSouce
	var destiProduct entities.ProductDestionation
	var isAdded bool
	for i := 0; i < count; i++ {
		productName := faker.Word()
		idProduct := uuid.NewString()
		sourceProduct.Id = idProduct
		destiProduct.Id = idProduct
		sourceProduct.ProductName = productName
		destiProduct.ProductName = productName
		sourceProduct.SellingPrice = float64(rand.Intn(140) + 10)
		sourceProduct.PromoPrice = float64(rand.Intn(140) + 7)
		sourceProduct.Qty = rand.Intn(14)
		destiProduct.Qty = 0
		destiProduct.SellingPrice = 0
		destiProduct.PromoPrice = 0

		txSource := dbSource.Begin()
		txDesti := dbDesti.Begin()

		if err := txSource.Debug().Create(&sourceProduct).Error; err != nil {
			txSource.Rollback()
			txDesti.Rollback()
			return err
		}

		if err := txDesti.Debug().Create(&destiProduct).Error; err != nil {
			txSource.Rollback()
			txDesti.Rollback()
			return err
		}

		txSource.Commit()
		txDesti.Commit()
		isAdded = true
	}

	if isAdded {
		err := redisConn.Set(ctx, "is_sync_product", "1", 0).Err()
		if err != nil {
			panic(err)
		}
	}

	return nil
}
