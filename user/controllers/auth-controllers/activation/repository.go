package activationAuth

import (
	"time"

	model "github.com/svaqqosov/k8s_microserices_starter/models"
	"gorm.io/gorm"
)

type Repository interface {
	ActivationRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryActivation(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ActivationRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	checkUserAccount := db.Debug().Select("*").Where("activation_code = ?", input.ActivationCode).Find(&users)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "ACTIVATION_NOT_FOUND_404"
		return &users, <-errorCode
	}

	if users.Active {
		errorCode <- "ACTIVATION_ACTIVE_400"
		return &users, <-errorCode
	}

	users.Active = input.Active
	users.UpdatedAt = time.Now().Local()
	users.ActivationCode = ""

	updateActivation := db.Debug().Update("active", true).Where("activation_code = ?", input.ActivationCode).Updates(&users)

	if updateActivation.Error != nil {
		errorCode <- "ACTIVATION_ACCOUNT_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
