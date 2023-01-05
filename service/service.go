package service

import (
	"api-product/command"
	"api-product/repository"
	"net/http"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type service struct {
	db        *gorm.DB
	repo      repository.IRepository
	response  command.Response
	redisConn *redis.Client
}

type IService interface {
	GetSourceProduct(w http.ResponseWriter, pagination *command.Pagination) command.Response
	GetDestinationProduct(w http.ResponseWriter, pagination *command.Pagination) command.Response
}

func NewService(db *gorm.DB, repo repository.IRepository, response command.Response, redisConn *redis.Client) IService {
	return &service{db, repo, response, redisConn}
}
