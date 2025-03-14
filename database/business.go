package database

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Business struct {
	GUID        uuid.UUID `gorm:"primaryKey" json:"guid"`
	BusinessUid string    `json:"business_uid"`
	NotifyUrl   string    `json:"notify_url"`
	Timestamp   uint64    `json:"timestamp"`
}

type BusinessDB interface {
	StoreBusiness(business *Business) error
}

type businessDB struct {
	gorm *gorm.DB
}

func NewBusinessDB(db *gorm.DB) BusinessDB {
	return &businessDB{gorm: db}
}

func (dao *businessDB) StoreBusiness(business *Business) error {
	err := dao.gorm.Table("businesses").Create(business).Error
	return err
}
