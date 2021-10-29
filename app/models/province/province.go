package province

import (
	"errors"
	"time"

	"github.com/gentildpinto/mangope-api/app/models/county"
	"github.com/gentildpinto/mangope-api/config/logger"
	"gorm.io/gorm"
)

type Province struct {
	ID             uint            `gorm:"primaryKey" json:"id"`
	Name           string          `gorm:"type:varchar(50);not null" json:"name"`
	Foundation     string          `gorm:"type:varchar(50)" json:"foundation"`
	Capital        string          `gorm:"type:varchar(50)" json:"capital"`
	Population     uint            `gorm:"type:int" json:"population"`
	Area           float64         `gorm:"type:float" json:"area"`
	PhonePrefix    string          `gorm:"type:varchar(4)" json:"phone_prefix"`
	GovernmentSite string          `gorm:"type:varchar(100)" json:"government_site"`
	Counties       []county.County `json:"counties"`
	CreatedAt      time.Time       `json:"-"`
	UpdatedAt      time.Time       `json:"-"`
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
