package database

import (
	"math/big"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Withdraws struct {
	GUID        uuid.UUID `gorm:"primaryKey" json:"guid"`
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
	Status      uint16    `json:"status"`
	Timestamp   uint64    `json:"timestamp"`
}

type WithdrawsDB interface {
	StoreWithdraw(*Withdraws) error
}

type withdrawsDB struct {
	gorm *gorm.DB
}

func NewWithdrawsDB(db *gorm.DB) WithdrawsDB {
	return &withdrawsDB{gorm: db}
}

func (dao *withdrawsDB) StoreWithdraw(withdraw *Withdraws) error {
	err := dao.gorm.Table("withdraws").Create(withdraw).Error
	return err
}