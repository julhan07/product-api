package repository

import (
	"api-product/command"
	"api-product/entities"

	"gorm.io/gorm"
)

type repository struct {
	dbSource *gorm.DB
	dbDesti  *gorm.DB
}

type IRepository interface {
	GetSourceProduct(pagination *command.Pagination) (res []entities.ProductSouce, count int64, err error)
	GetDestinationProduct(pagination *command.Pagination) (res []entities.ProductDestionation, count int64, err error)
	UpdateDestinationProduct(tx *gorm.DB, req entities.ProductDestionation) (err error)
}

func NewRepository(dbSource, dbDesti *gorm.DB) IRepository {
	return &repository{dbSource, dbDesti}
}
