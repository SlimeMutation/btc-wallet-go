package database

import (
	"math/big"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vins struct {
	GUID             uuid.UUID `gorm:"primaryKey" json:"guid"`
	Address          string    `json:"address"`
	Txid             string    `json:"txid"`
	Vout             uint8     `json:"vout"`
	Script           string    `json:"script"`
	Witness          string    `json:"witness"`
	Amount           *big.Int  `gorm:"serializer:u256"`
	SpendTxHash      string    `json:"spend_tx_hash"`
	SpendBlockHeight *big.Int  `gorm:"serializer:u256"`
	IsSpend          bool      `json:"is_spend"`
	Timestamp        uint64    `json:"timestamp"`
}

type VinsDB interface {
	StoreVin(*Vins) error
}

type vinsDB struct {
	gorm *gorm.DB
}

func NewVinsDB(db *gorm.DB) VinsDB {
	return &vinsDB{gorm: db}
}

func (dao *vinsDB) StoreVin(vin *Vins) error {
	err := dao.gorm.Table("vins").Create(vin).Error
	return err
}
