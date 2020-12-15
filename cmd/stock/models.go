package stock

import "time"

type ProductData struct {
	ID int `json:"id" gorm:"primary_key"`
	Date time.Time `json:"date"`
	Salon string `json:"salon"`
	Product string `json:"product"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
	Supplier string `json:"supplier"`
	Brand string `json:"brand"`
	Category string `json:"category"`
	SubBrand string `json:"sub_brand"`
	Type string `json:"type"`
}

type Total struct {
	S string `json:"salon"`
	T int `json:"total"`
}