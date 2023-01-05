package entities

import (
	"time"

	"gorm.io/gorm"
)

type ProductDestionation struct {
	ProductEntity
}

func (ProductDestionation) TableName() string {
	return "destination_product"
}

func (p *ProductDestionation) BeforeCreate(tx *gorm.DB) (err error) {
	tNow := time.Now()
	p.CreatedAt = &tNow
	return
}
