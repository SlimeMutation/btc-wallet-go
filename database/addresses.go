package database

import (
	"math/big"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Addresses struct {
	GUID        uuid.UUID `gorm:"primaryKey" json:"guid"`
	Address     string    `json:"address"`
	AddressType uint8     `json:"address_type"` //0:用户地址；1:热钱包地址(归集地址)；2:冷钱包地址
	PublicKey   *big.Int  `json:"public_key"`
	Timestamp   uint64    `json:"timestamp"`
}

type AddressesDB interface {
	StoreAddress(*Addresses) error
}

type addressesDB struct {
	gorm *gorm.DB
}

func NewAddressesDB(db *gorm.DB) AddressesDB {
	return &addressesDB{gorm: db}
}

func (dao *addressesDB) StoreAddress(address *Addresses) error {
	err := dao.gorm.Table("addresses").Create(address).Error
	return err
}
