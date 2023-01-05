package service

import (
	"api-product/command"
	"api-product/entities"
	"context"
	"log"
	"net/http"
	"time"
)

var ctx = context.Background()

func (s *service) GetSourceProduct(w http.ResponseWriter, pagination *command.Pagination) command.Response {
	res, count, err := s.repo.GetSourceProduct(pagination)
	if err != nil {
		return s.response.BadRequest(err.Error())
	}

	if s.checkSyncData() {
		go func() {
			dp, err := s.syncListDestinationProduct()
			if err != nil {
				log.Println("error sycn product destionation", err.Error())
			}

			if err := s.syncDestinationProduct(dp); err != nil {
				log.Println("error sycn product destionation", err.Error())
			}
		}()
	}

	return s.response.Success(pagination, res, count)
}

func (s *service) GetDestinationProduct(w http.ResponseWriter, pagination *command.Pagination) command.Response {
	res, count, err := s.repo.GetDestinationProduct(pagination)
	if err != nil {
		return s.response.BadRequest(err.Error())
	}

	return s.response.Success(pagination, res, count)
}

func (s *service) syncListDestinationProduct() (dp []entities.ProductDestionation, err error) {

	pagination := command.Pagination{IsPagination: false}

	ps, _, err := s.repo.GetSourceProduct(&pagination)
	if err != nil {
		return nil, err
	}

	res, _, err := s.repo.GetDestinationProduct(&pagination)
	if err != nil {
		return nil, err
	}

	t := time.Now()
	for _, v := range res {
		for _, sv := range ps {
			if v.Id == sv.Id {
				v.Qty = sv.Qty
				v.PromoPrice = sv.PromoPrice
				v.SellingPrice = sv.SellingPrice
				v.UpdatedAt = &t
				dp = append(dp, v)
			}
		}
	}

	return dp, nil
}

func (s *service) syncDestinationProduct(dp []entities.ProductDestionation) (err error) {
	var isUpdated bool
	for _, v := range dp {
		tx := s.db.Begin()
		if err := s.repo.UpdateDestinationProduct(tx, v); err != nil {
			tx.Rollback()
			return err
		}

		tx.Commit()
		isUpdated = true
	}

	if isUpdated {
		if err := s.redisConn.Set(ctx, "is_sync_product", "0", 0).Err(); err != nil {
			log.Println(err.Error())
		}
	}
	return nil
}

func (s *service) checkSyncData() bool {
	val, err := s.redisConn.Get(ctx, "is_sync_product").Result()
	if err != nil {
		log.Println(err.Error())
	}

	if val == "1" {
		return true
	}
	return false
}
