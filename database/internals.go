package database

import (
	"math/big"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Internals struct {
	GUID        uuid.UUID `gorm:"primaryKey" json:"guid"`
	Timestamp   uint64    `json:"timestamp"`
	Status      uint16    `json:"status"`
	BlockHash   string    `json:"block_hash"`
	BlockNumber *big.Int  `gorm:"serializer:u256" json:"block_number"`
	Hash        string    `json:"hash"`
	FromAddress string    `json:"from_address"`
	ToAddress   string    `json:"to_address"`
	Amount      string    `json:"amount"`
	Fee         string    `json:"fee"`
	LockTime    *big.Int  `gorm:"serializer:u256" json:"lock_time"`
	Version     string    `json:"version"`
	TxSignHex   string    `json:"tx_sign_hex"`
}

type InternalsDB interface {
	StoreInternal(*Internals) error
}

type internalsDB struct {
	gorm *gorm.DB
}

func NewInternalsDB(db *gorm.DB) InternalsDB {
	return &internalsDB{gorm: db}
}

func (dao *internalsDB) StoreInternal(internal *Internals) error {
	err := dao.gorm.Table("internals").Create(internal).Error
	return err
}