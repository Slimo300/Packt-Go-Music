package dblayer

import (
	"errors"

	"github.com/Slimo300/Packt-Go-Music/backend/src/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con)
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

func (db *DBORM) GetCustomerByName(firstname, lastname string) (customer models.Customer, err error) {
	return customer, db.Where(&models.Customer{FirstName: firstname, LastName: lastname}).Find(&customer).Error
}

func (db *DBORM) GetCustomerById(id int) (customer models.Customer, err error) {
	return customer, db.First(&customer, id).Error
}

func (db *DBORM) GetProduct(id int) (product models.Product, err error) {
	return product, db.First(&product, id).Error
}

func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {
	hashPassword(&customer.Pass)
	customer.LoggedIn = true
	return customer, db.Create(&customer).Error
}

func (db *DBORM) SignInUser(email, pass string) (customer models.Customer, err error) {
	result := db.Table("customers").Where(&models.Customer{Email: email})
	err = result.First(&customer).Error
	if err != nil {
		return customer, err
	}
	if !checkPassword(customer.Pass, pass) {
		return customer, ErrINVALIDPASSWORD
	}
	customer.Pass = ""
	err = result.Update("loggedin", 1).Error
	if err != nil {
		return customer, err
	}

	return customer, result.Find(&customer).Error
}

func (db *DBORM) SignOutUserById(id int) error {
	customer := models.Customer{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	return db.Table("customers").Where(&customer).Update("loggedin", 0).Error
}

func (db *DBORM) GetCustomerOrdersById(id int) (orders []models.Order, err error) {
	return orders, db.Table("orders").Select("*").
		Joins("join customers on customers.id = customer_id").
		Joins("join products on product.id = product_id").
		Where("customer_id=?", id).Scan(&orders).Error
}

func (db *DBORM) AddOrder(order models.Order) error {
	return db.Create(&order).Error
}

func (db *DBORM) GetCreditCardCID(id int) (string, error) {
	customerWithCCID := struct {
		models.Customer
		CCID string `gorm:"column:cc_customerid"`
	}{}
	return customerWithCCID.CCID, db.First(&customerWithCCID, id).Error
}

func (db *DBORM) SaveCreditCardForCustomer(id int, ccid string) error {
	result := db.Table("customers").Where("id=?", id)
	return result.Update("cc_customerid", ccid).Error
}

func hashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}

	// convert password string to byte slice so that we can use it with the bcrypt package
	sBytes := []byte(*s)

	// Obtian hashed password via the GenerateFromPassword() method
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update password string with the hashed version
	*s = string(hashedBytes)
	return nil
}

func checkPassword(existingHash, incomingPass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(incomingPass)) == nil
}
