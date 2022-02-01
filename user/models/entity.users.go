package model

import (
	"time"

	"github.com/google/uuid"
	util "github.com/svaqqosov/k8s_microserices_starter/utils"
	"gorm.io/gorm"
)

type EntityUsers struct {
	ID             string `gorm:"primaryKey;"`
	Fullname       string `gorm:"type:varchar(255);unique;not null"`
	Email          string `gorm:"type:varchar(255);unique;not null"`
	Password       string `gorm:"type:varchar(255);not null"`
	Active         bool   `gorm:"type:bool;default:false"`
	ActivationCode string `gorm:"type:varchar(36);not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (entity *EntityUsers) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.Password = util.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	entity.ActivationCode = uuid.NewString()
	return nil
}

func (entity *EntityUsers) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
