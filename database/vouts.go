package database

import (
	"math/big"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vouts struct {
	GUID      uuid.UUID `gorm:"primaryKey" json:"guid"`
	Address   string    `json:"address"`
	N         uint16    `json:"n"`
	Script    string    `json:"script"`
	Amount    *big.Int  `gorm:"serializer:u256" json:"amount"`
	Timestamp uint64    `json:"timestamp"`
}

type VoutsDB interface {
	StoreVout(*Vouts) error
}

type voutsDB struct {
	gorm *gorm.DB
}

func NewVoutsDB(db *gorm.DB) VoutsDB {
	return &voutsDB{gorm: db}
}

func (dao *voutsDB) StoreVout(vout *Vouts) error {
	err := dao.gorm.Table("vouts").Create(vout).Error
	return err
}