package county

import (
	"errors"
	"time"

	"github.com/gentildpinto/mangope-api/config/logger"
	"gorm.io/gorm"
)

type (
	Tabler interface {
		TableName() string
	}

	County struct {
		ID         uint      `gorm:"primaryKey" json:"id"`
		Name       string    `gorm:"type:varchar(50);not null" json:"name"`
		ProvinceID uint      `gorm:"not null" json:"province_id"`
		CreatedAt  time.Time `json:"-"`
		UpdatedAt  time.Time `json:"-"`
	}
)

func (County) TableName() string {
	return "counties"
}

var db *gorm.DB

func Initialize(database *gorm.DB) (err error) {
	if database == nil {
		return errors.New("invalid database")
	}
	db = database

	return
}

func GetAll() (counties []County, err error) {
	if err = logger.Log(db.Find(&counties).Error); err != nil {
		return
	}
	return
}

func GetByID(id int) (county County, err error) {
	if err = logger.Log(db.First(&county, id).Error); err != nil {
		return
	}
	return
}
