package loginAuth

import (
	model "github.com/firmanJS/store-app/models"
	util "github.com/firmanJS/store-app/utils"
	"gorm.io/gorm"
)

type Repository interface {
	LoginRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryLogin(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) LoginRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Username = input.Username
	users.Password = input.Password

	checkUserAccount := db.Debug().Select("*").Where("username = ?", input.Username).Find(&users)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- util.NOT_FOUND
		return &users, <-errorCode
	}

	comparePassword := util.ComparePassword(users.Password, input.Password)

	if comparePassword != nil {
		errorCode <- util.FAILED
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
