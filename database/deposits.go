package database

import (
	"math/big"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Deposits struct {
	GUID        uuid.UUID `gorm:"primaryKey" json:"guid"`
	BlockHash   string    `json:"block_hash"`
	BlockNumber *big.Int  `gorm:"serializer:u256" json:"block_number"`
	Hash        string    `json:"hash"`
	FromAddress string    `json:"from_address"`
	ToAddress   string    `json:"to_address"`
	Fee         *big.Int  `gorm:"serializer:u256" json:"fee"`
	LockTime    *big.Int  `gorm:"serializer:u256" json:"lock_time"`
	Version     string    `json:"version"`
	Amount      *big.Int  `gorm:"serializer:u256" json:"amount"`
	Confirms    uint16    `json:"confirms"`
	Status      uint16    `json:"status"`
	Timestamp   uint64    `json:"timestamp"`
}

type DepositsDB interface {
	StoreDeposit(*Deposits) error
}

type depositsDB struct {
	gorm *gorm.DB
}

func NewDepositsDB(db *gorm.DB) DepositsDB {
	return &depositsDB{gorm: db}
}

func (dao *depositsDB) StoreDeposit(deposit *Deposits) error {
	err := dao.gorm.Table("deposits").Create(deposit).Error
	return err
}
