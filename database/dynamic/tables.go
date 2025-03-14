package dynamic

import (
	"fmt"

	"github.com/SlimeMutation/btc-wallet-go/database"
)

func CreateTableFromTemplate(requestId string, db *database.DB) {
	createAddresses(requestId, db)
}

func createAddresses(requestId string, db *database.DB) {
	tableName := "addresses"
	tableNameByChainId := fmt.Sprintf("addresses_%s", requestId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}
