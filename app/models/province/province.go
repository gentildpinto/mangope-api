package province

import (
	"errors"
	"time"

	"github.com/gentildpinto/mangope-api/config/logger"
	"gorm.io/gorm"
)

type Province struct {
	ID             uint      `gorm:"primaryKey"`
	Name           string    `gorm:"type:varchar(50);not null"`
	Foundation     string    `gorm:"type:varchar(50)"`
	Capital        string    `gorm:"type:varchar(50)"`
	Papulation     int       `gorm:"type:int"`
	Area           float64   `gorm:"type:float"`
	PhonePrefix    string    `gorm:"type:varchar(4)"`
	GovernmentSite string    `gorm:"type:varchar(100)"`
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

func GetAllProvinces() (provinces []Province, err error) {
	if err = logger.Log(db.Find(&provinces).Error); err != nil {
		return
	}
	return
}
