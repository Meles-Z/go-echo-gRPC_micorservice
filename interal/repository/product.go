package repository

import (
	"context"
	"errors"

	"github.com/meles-z/go-grpc-microsevice/interal/entities"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
	"gorm.io/gorm"
)

type ProductRepoImp struct {
	DB *gorm.DB
	*order.UnimplementedProductServiceServer
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, req *order.CreateProductRequest) (*order.CreateProductResponse, error)
	GetAllProducts(ctx context.Context, req *order.GetAllProductsRequest) (*order.GetAllProductsResponse, error)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepoImp{DB: db}
}

func (s *ProductRepoImp) CreateProduct(ctx context.Context, req *order.CreateProductRequest) (*order.CreateProductResponse, error) {
	prod := req.GetProduct()
	data := entities.Product{
		Name:        prod.Name,
		Description: prod.Description,
		Price:       prod.Price,
		StockQty:    int(prod.StockQuantity),
	}
	if err := s.DB.Create(&data).Error; err != nil {
		return nil, errors.New("failed to create product:" + err.Error())
	}
	return &order.CreateProductResponse{
		Product: &order.Product{
			Name:          prod.Name,
			Description:   prod.Description,
			Price:         prod.Price,
			StockQuantity: prod.StockQuantity,
		},
	}, nil
}

func (s *ProductRepoImp) GetAllProducts(ctx context.Context, req *order.GetAllProductsRequest) (*order.GetAllProductsResponse, error) {
	var product []entities.Product
	if err := s.DB.Find(&product).Error; err != nil {
		return nil, errors.New("Error to find record:" + err.Error())
	}
	var prod []*order.Product
	for _, pro := range product {
		prod = append(prod, &order.Product{
			Id:            pro.ID,
			Name:          pro.Name,
			Description:   pro.Description,
			Price:         pro.Price,
			StockQuantity: int64(pro.StockQty),
		})
	}
	return &order.GetAllProductsResponse{
		Product: prod,
	}, nil
}
