package database

import (
	"math/big"

	"gorm.io/gorm"
)

type Blocks struct {
	Hash      string   `gorm:"primaryKey" json:"hash"`
	PrevHash  string   `json:"prev_hash"`
	Number    *big.Int `gorm:"serializer:u256"`
	Timestamp uint64   `json:"timestamp"`
}

type BlocksDB interface {
	StoreBlock(*Blocks) error
}

type blocksDB struct {
	gorm *gorm.DB
}

func NewBlocksDB(db *gorm.DB) BlocksDB {
	return &blocksDB{gorm: db}
}

func (dao *blocksDB) StoreBlock(block *Blocks) error {
	err := dao.gorm.Table("blocks").Create(block).Error
	return err
}
