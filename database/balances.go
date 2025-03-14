package database

import (
	"math/big"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Balances struct {
	GUID        uuid.UUID `gorm:"primaryKey" json:"guid"`
	Address     string    `json:"address"`
	AddressType uint8     `json:"address_type"` //0:用户地址；1:热钱包地址(归集地址)；2:冷钱包地址
	Balance     *big.Int  `gorm:"serializer:u256"`
	LockBalance *big.Int  `gorm:"serializer:u256"`
	Timestamp   uint64    `json:"timestamp"`
}

type BalancesDB interface {
	StoreBalance(*Balances) error
}

type balancesDB struct {
	gorm *gorm.DB
}

func NewBalancesDB(db *gorm.DB) BalancesDB {
	return &balancesDB{gorm: db}
}

func (dao *balancesDB) StoreBalance(balance *Balances) error {
	err := dao.gorm.Table("balances").Create(balance).Error
	return err
}
