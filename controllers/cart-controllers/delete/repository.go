package deleteCart

import (
	model "github.com/firmanJS/store-app/models"
	util "github.com/firmanJS/store-app/utils"
	"gorm.io/gorm"
)

type Repository interface {
	DeletedCartRepository(input *model.EntityCart) (*model.EntityCart, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryDelete(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) DeletedCartRepository(input *model.EntityCart) (*model.EntityCart, string) {

	var dCarts model.EntityCart
	db := r.db.Model(&dCarts)
	errorCode := make(chan string, 1)

	checkdCartId := db.Debug().Select("*").Where("id = ?", input.Id).Find(&dCarts)

	if checkdCartId.RowsAffected < 1 {
		errorCode <- util.NOT_FOUND
		return &dCarts, <-errorCode
	}

	deletedCartId := db.Debug().Select("id").Where("id = ?", input.Id).Find(&dCarts).Delete(&dCarts)

	if deletedCartId.Error != nil {
		errorCode <- util.FAILED
		return &dCarts, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &dCarts, <-errorCode
}
