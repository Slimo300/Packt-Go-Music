package dblayer

import (
	"errors"

	"github.com/Slimo300/Packt-Go-Music/backend/src/models"
)

type DBLayer interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerById(int) (models.Customer, error)
	GetProduct(uint) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(models.Customer) (models.Customer, error)
	SignOutUserById(int) error
	GetCustomerOrdersById(int) ([]models.Order, error)
	AddOrder(models.Order) error
	GetCreditCatdCID(int) (string, error)
	SaveCreditCardForCustomer(int, string) error
}

var ErrINVALIDPASSWORD = errors.New("Invalid password")
