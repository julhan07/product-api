package app

import (
	"api-product/app/handler"
	"api-product/command"
	"api-product/config"
	"api-product/repository"
	"api-product/service"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func RunApp(conf *config.App, dbSource, dbDesti *gorm.DB, redisConn *redis.Client) {
	resp := command.NewReponse()

	repo := repository.NewRepository(dbSource, dbDesti)
	service := service.NewService(dbDesti, repo, resp, redisConn)
	handler := handler.NewHandler(service, resp)

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router(handler), // Pass our instance of gorilla/mux in.
	}

	fmt.Printf("server runnning on http://%s:%s", conf.Host, conf.Port)
	fmt.Println("")
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
