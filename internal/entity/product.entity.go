package entity

import (
	"mime/multipart"
	"time"
)

type ProductData struct {
	ProductID                  int       `json:"product_id" binding:"required"`
	ProductTitle               string    `json:"product_title" binding:"required"`
	ProductDescription         string    `json:"product_description" binding:"required"`
	ProductPrice               int       `json:"product_price" binding:"required"`
	ProdcutCreatedByEmployeeID int       `json:"created_by_employee_id" binding:"required"`
	ProductCoverPath           string    `json:"product_cover" binding:"required"`
	CreatedAt                  time.Time `json:"created_at" binding:"required"`
}

type ProductType struct {
	ProductID                  int    `json:"product_id" binding:"required"`
	ProductTitle               string `json:"product_title" binding:"required"`
	ProductDescription         string `json:"product_description" binding:"required"`
	ProductPrice               int    `json:"product_price" binding:"required"`
	ProdcutCreatedByEmployeeID string `json:"created_by_employee_id" binding:"required"`
	ProductCover               string `json:"product_cover" binding:"required"`
}

type ProductPayload struct {
	ProductTitle               string                `form:"product_title" binding:"required"`
	ProductDescription         string                `form:"product_description" binding:"required"`
	ProductPrice               int64                 `form:"product_price" binding:"required"`
	ProdcutCreatedByEmployeeID int                   `form:"created_by_employee_id"`
	ProductCover               *multipart.FileHeader `form:"-"` //  The form:"-" tag tells Fiber to skip trying to decode the file automatically — you'll assign it manually.
	ProductCoverPath           string
}

type ProductPayloadUpdate struct {
	ProductID                  int                   `form:"product_id" binding:"required"`
	ProductTitle               string                `form:"product_title" binding:"required"`
	ProductDescription         string                `form:"product_description" binding:"required"`
	ProductPrice               int                   `form:"product_price" binding:"required"`
	ProdcutCreatedByEmployeeID int                   `form:"created_by_employee_id"`
	ProductCover               *multipart.FileHeader `form:"-"` //  The form:"-" tag tells Fiber to skip trying to decode the file automatically — you'll assign it manually.
	ProductCoverPath           string
}
type InputProductType struct {
	EmployeeID int `json:"employee_id" binding:"required"`
	ProductID  int `json:"product_id" binding:"required"`
}
