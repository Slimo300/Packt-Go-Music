package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Image       string  `json:"img"`
	ImagAlt     string  `gorm:"column:imgalt" json:"imgalt"`
	Price       float32 `json:"price"`
	Promotion   float32 `json:"promotion"`
	ProductName string  `gorm:"column:productname" json:"productname"`
	Description string  `json:"desc"`
}

func (Product) TableName() string {
	return "products"
}

type Customer struct {
	gorm.Model
	Name      string  `json:"name"`
	FirstName string  `gorm:"column:firstname" json:"firstname"`
	LastName  string  `gorm:"column:lastname" json:"lastname"`
	Email     string  `gorm:"column:email" json:"email"`
	Pass      string  `json:"password"`
	LoggedIn  bool    `gorm:"column:loggedin" json:"loggedin"`
	Orders    []Order `json:"orders"`
}

func (Customer) TableName() string {
	return "customers"
}

type Order struct {
	gorm.Model
	Product
	Customer
	CustomerID   int       `gorm:"column:customerid" json:"customer_id"`
	ProductID    int       `gorm:"column:productid" json:"product_id"`
	Price        float32   `gorm:"column:price" json:"sell_price"`
	PurchaseDate time.Time `gorm:"column:purchase_date" json:"purchase_date"`
}

func (Order) TableName() string {
	return "orders"
}
