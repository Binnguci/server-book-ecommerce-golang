package repository

import (
	"math"
	"server-furniture-ecommerce-gin/global"
	"server-furniture-ecommerce-gin/internal/model"
)

type IProductRepository interface {
	FindAll(page, limit int) ([]model.Product, int, error)
	FindByProduct(productName string) ([]model.Product, int, error)
}
