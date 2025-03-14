package database

import (
	"math/big"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transactions struct {
	GUID        uuid.UUID `gorm:"primaryKey" json:"guid"`
	BlockHash   string    `json:"block_hash"`
	BlockNumber *big.Int  `gorm:"serializer:u256" json:"block_number"`
	Hash        string    `json:"hash"`
	FromAddress string    `json:"from_address"`
	ToAddress   string    `json:"to_address"`
	Fee         *big.Int  `gorm:"serializer:u256" json:"fee"`
	LockTime    *big.Int  `gorm:"serializer:u256" json:"lock_time"`
	Version     string    `json:"version"`
	Amount      string    `json:"amount"`
	Status      uint16    `json:"status"`
	TxType      string    `json:"tx_type"`
	Timestamp   uint64    `json:"timestamp"`
}

type TransactionsDB interface {
	StoreTransaction(*Transactions) error
}

type transactionsDB struct {
	gorm *gorm.DB
}

func NewTransactionsDB(db *gorm.DB) TransactionsDB {
	return &transactionsDB{gorm: db}
}

func (dao *transactionsDB) StoreTransaction(transaction *Transactions) error {
	err := dao.gorm.Table("transactions").Create(transaction).Error
	return err
}
