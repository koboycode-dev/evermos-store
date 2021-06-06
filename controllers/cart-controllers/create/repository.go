package createCart

import (
	model "github.com/firmanJS/store-app/models"
	util "github.com/firmanJS/store-app/utils"
	"gorm.io/gorm"
)

type Repository interface {
	CreateCartRepository(input *model.EntityCart) (*model.EntityCart, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateCartRepository(input *model.EntityCart) (*model.EntityCart, string) {

	var Carts model.EntityCart
	db := r.db.Model(&Carts)
	errorCode := make(chan string, 1)

	Carts.Order_Id = input.Order_Id
	Carts.Product_Id = input.Product_Id
	Carts.Quantity = input.Quantity
	Carts.Note = input.Note

	addNewCart := db.Debug().Create(&Carts)
	db.Commit()

	if addNewCart.Error != nil {
		errorCode <- util.FAILED
		return &Carts, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &Carts, <-errorCode
}
