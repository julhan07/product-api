package repository

import (
	"api-product/command"
	"api-product/entities"
	"fmt"

	"gorm.io/gorm"
)

func (r *repository) GetSourceProduct(pagination *command.Pagination) (res []entities.ProductSouce, count int64, err error) {

	order := fmt.Sprintf("%s %s", pagination.SortBy, pagination.Order)
	keyword := `%` + pagination.Keyword + `%`

	query := r.dbSource.Debug()

	if pagination.IsPagination {
		if pagination.Keyword != "" {
			query = query.Where("product_name ILIKE ?", keyword)
		}

		query = query.Order(order).Limit(pagination.Limit).Offset(pagination.Offset)
	}

	if err = query.Find(&res).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Model(&res).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return res, count, nil
}

func (r *repository) GetDestinationProduct(pagination *command.Pagination) (res []entities.ProductDestionation, count int64, err error) {

	order := fmt.Sprintf("%s %s", pagination.SortBy, pagination.Order)
	keyword := `%` + pagination.Keyword + `%`

	query := r.dbDesti.Debug()

	if pagination.IsPagination {
		if pagination.Keyword != "" {
			query = query.Where("product_name ILIKE ?", keyword, keyword, keyword, keyword)
		}

		query = query.Order(order).Limit(pagination.Limit).Offset(pagination.Offset)
	}

	if err = query.Find(&res).Error; err != nil {
		return nil, 0, err
	}

	if err = query.Model(&res).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return res, count, nil
}

func (r *repository) UpdateDestinationProduct(tx *gorm.DB, req entities.ProductDestionation) (err error) {

	if err = tx.Debug().Save(req).Error; err != nil {
		return err
	}
	return nil
}
