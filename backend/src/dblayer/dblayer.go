package dblayer

import (
	"github.com/Slimo300/Packt-Go-Music/backend/src/models"
)

type DBLayer interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerById(int) (models.Customer, error)
	GetProduct(uint) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(username, password string)
	SignOutUserById(int) error
	GetCustomerOrdersById(int) ([]models.Order, error)
}
