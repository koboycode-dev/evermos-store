package updateCart

import (
	model "github.com/firmanJS/store-app/models"
	util "github.com/firmanJS/store-app/utils"
	"gorm.io/gorm"
)

type Repository interface {
	UpdateCartRepository(input *model.EntityCart) (*model.EntityCart, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateCartRepository(input *model.EntityCart) (*model.EntityCart, string) {

	var Carts model.EntityCart
	db := r.db.Model(&Carts)
	errorCode := make(chan string, 1)

	Carts.Id = input.Id

	checkCartId := db.Debug().Select("id").Where("id = ?", input.Id).Find(&Carts)

	if checkCartId.RowsAffected < 1 {
		errorCode <- util.NOT_FOUND
		return &Carts, <-errorCode
	}

	Carts.Product_Id = input.Product_Id
	Carts.Quantity = input.Quantity
	Carts.Note = input.Note

	updateCart := db.Debug().Select("product_id", "quantity", "note", "updated_at").Where("id = ?", input.Id).Updates(Carts)

	if updateCart.Error != nil {
		errorCode <- util.FAILED
		return &Carts, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &Carts, <-errorCode
}
