package province

import (
	"errors"
	"time"

	"github.com/gentildpinto/mangope-api/app/models/county"
	"github.com/gentildpinto/mangope-api/config/logger"
	"gorm.io/gorm"
)

type Province struct {
	ID             uint    `gorm:"primaryKey"`
	Name           string  `gorm:"type:varchar(50);not null"`
	Foundation     string  `gorm:"type:varchar(50)"`
	Capital        string  `gorm:"type:varchar(50)"`
	Papulation     uint    `gorm:"type:int"`
	Area           float64 `gorm:"type:float"`
	PhonePrefix    string  `gorm:"type:varchar(4)"`
	GovernmentSite string  `gorm:"type:varchar(100)"`
	Counties       []county.County
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}

var db *gorm.DB

func Initialize(database *gorm.DB) (err error) {
	if database == nil {
		return errors.New("invalid database")
	}
	db = database

	return
}

func GetAll() (provinces []Province, err error) {
	if err = logger.Log(db.Preload("Counties").Find(&provinces).Error); err != nil {
		return
	}
	return
}

func GetByID(id int) (province Province, err error) {
	if err = logger.Log(db.Preload("Counties").First(&province, id).Error); err != nil {
		return
	}
	return
}
