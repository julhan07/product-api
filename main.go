package main

import (
	"api-product/app"
	"api-product/config"
	"api-product/entities"
	"fmt"
	"log"
	"os"

	"api-product/infra"

	"github.com/joho/godotenv"
)

func main() {

	//  load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get all configuration
	conf := config.GetConfiguration()

	// redis connectin
	redisConn := infra.GetRedisConnection(&conf.RedisConfiguration)

	// get psql connection
	dbSource := infra.GetConnection(&conf.PsqSource)
	dbDesti := infra.GetConnection(&conf.PsqDestination)
	if conf.App.Env != "production" {
		dbSource.Debug().AutoMigrate(&entities.ProductSouce{})
		dbDesti.Debug().AutoMigrate(&entities.ProductDestionation{})
	}

	// auto seed table of db
	if os.Getenv("DB_SEED") == "1" {
		go func() {
			if err := infra.RunSeed(dbSource, dbDesti, 500, redisConn); err != nil {
				fmt.Println("error seed data :", err.Error())
			}
		}()
	}

	// run app
	app.RunApp(&conf.App, dbSource, dbDesti, redisConn)
}
