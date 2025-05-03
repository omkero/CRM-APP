package types

import "time"

type UserData struct {
	Username string `json:"username" binding:"required"`
}

type UserModel struct {
	CustomerName    string `json:"customer_name" binding:"required"`
	CustomerRole    string `json:"customer_role" binding:"required"`
	PreferedProduct string `json:"prefered_product" binding:"required"`
}

type CustomerModel struct {
	CustomerID      int       `json:"customer_id" binding:"required"`
	CustomerName    string    `json:"customer_name" binding:"required"`
	CustomerRole    string    `json:"customer_role" binding:"required"`
	PreferedProduct string    `json:"prefered_product" binding:"required"`
	CreatedAt       time.Time `json:"created_at" binding:"required"`
}
