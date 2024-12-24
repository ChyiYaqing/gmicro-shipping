package db

import (
	"context"
	"fmt"

	"github.com/chyiyaqing/gmicro-shipping/internal/application/core/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(sqliteDB string) (*Adapter, error) {
	db, openErr := gorm.Open(sqlite.Open(sqliteDB), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db open %v error: %v", sqliteDB, openErr)
	}
	err := db.AutoMigrate(&Shipping{})
	if err != nil {
		return nil, fmt.Errorf("db migrate error: %v", err)
	}
	return &Adapter{db: db}, nil
}

type Shipping struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderID    int64
	Address    string
}

func (a Adapter) Get(ctx context.Context, id string) (domain.Shipping, error) {
	var shippingEntity Shipping
	res := a.db.WithContext(ctx).First(&shippingEntity, id)
	shipping := domain.Shipping{
		ID:         int64(shippingEntity.ID),
		CustomerID: shippingEntity.CustomerID,
		Status:     shippingEntity.Status,
		OrderId:    shippingEntity.OrderID,
		CreatedAt:  shippingEntity.CreatedAt.UnixNano(),
		Address:    shippingEntity.Address,
	}
	return shipping, res.Error
}

func (a Adapter) Save(ctx context.Context, shipping *domain.Shipping) error {
	shippingModel := Shipping{
		CustomerID: shipping.CustomerID,
		Status:     shipping.Status,
		OrderID:    shipping.OrderId,
		Address:    shipping.Address,
	}
	res := a.db.WithContext(ctx).Create(&shippingModel)
	if res.Error == nil {
		shipping.ID = int64(shippingModel.ID)
	}
	return res.Error
}
