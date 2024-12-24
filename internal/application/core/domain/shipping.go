package domain

import "time"

type Shipping struct {
	ID         int64  `json:"id"`
	CustomerID int64  `json:"customer_id"`
	Status     string `json:"status"`
	OrderId    int64  `json:"order_id"`
	Address    string `json:"address"`
	CreatedAt  int64  `json:"created_at"`
}

func NewShipping(customerId int64, orderId int64, address string) Shipping {
	return Shipping{
		CreatedAt:  time.Now().Unix(),
		Status:     "Pending",
		CustomerID: customerId,
		OrderId:    orderId,
		Address:    address,
	}
}
