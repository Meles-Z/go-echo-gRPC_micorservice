package repository

import (
	"context"
	"errors"

	"github.com/meles-z/go-grpc-microsevice/interal/entities"
	order "github.com/meles-z/go-grpc-microsevice/pkg/pb"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	CreateOrderItem(ctx context.Context, req *order.CreateOrderItemRequest) (*order.CreateOrderItemResponse, error)
	GetAllOrdersItem(ctx context.Context, req *order.GetAllOrderItemsRequest, stream order.OrderItemService_GetAllOrdersItemServer) error
}

type OrderItemImp struct {
	DB *gorm.DB
	*order.UnimplementedOrderItemServiceServer
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &OrderItemImp{
		DB: db,
	}
}

func (s *OrderItemImp) CreateOrderItem(ctx context.Context, req *order.CreateOrderItemRequest) (*order.CreateOrderItemResponse, error) {
	item := req.GetOrderItem()
	newItem := entities.OrderItem{
		OrderID:    item.OrderId,
		ProductID:  item.ProductId,
		Quantity:   int(item.Quantity),
		UnitPrice:  item.UnitPrice,
		TotalPrice: item.TotalPrice,
	}

	err := s.DB.Create(&newItem)
	if err.RowsAffected == 0 {
		return nil, errors.New("error to order item:" + err.Error.Error())
	}
	return &order.CreateOrderItemResponse{
		OrderItem: &order.OrderItem{
			OrderId:    item.OrderId,
			ProductId:  item.ProductId,
			Quantity:   item.Quantity,
			UnitPrice:  item.UnitPrice,
			TotalPrice: item.TotalPrice,
		},
	}, nil
}

func (s *OrderItemImp) GetAllOrdersItem(ctx context.Context, req *order.GetAllOrderItemsRequest, stream order.OrderItemService_GetAllOrdersItemServer) error {
	var dbUser []entities.OrderItem
	if err := s.DB.Find(&dbUser).Error; err != nil {
		return errors.New("failed to return all items:" + err.Error())
	}
	for _, item := range dbUser {
		grpcItem := &order.OrderItem{
			Id:         item.ID,
			OrderId:    item.OrderID,
			ProductId:  item.ProductID,
			Quantity:   int32(item.Quantity),
			UnitPrice:  item.UnitPrice,
			TotalPrice: item.TotalPrice,
		}
		if err := stream.Send(grpcItem); err != nil {
			return err
		}
	}
	return nil
}
