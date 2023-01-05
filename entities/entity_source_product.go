package entities

import (
	"time"

	"gorm.io/gorm"
)

type ProductSouce struct {
	ProductEntity
}

func (p *ProductSouce) BeforeCreate(tx *gorm.DB) (err error) {
	tNow := time.Now()
	p.CreatedAt = &tNow
	return
}

func (ProductSouce) TableName() string {
	return "source_product"
}
