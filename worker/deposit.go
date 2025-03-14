package worker

import (
	"context"

	"github.com/SlimeMutation/btc-wallet-go/common/tasks"
	"github.com/SlimeMutation/btc-wallet-go/config"
	"github.com/SlimeMutation/btc-wallet-go/database"
	"github.com/SlimeMutation/btc-wallet-go/rpcclient"
)

type Deposit struct {
	BaseSynchronizer

	confirms       uint8
	latestHeader   rpcclient.BlockHeader
	resourceCtx    context.Context
	resourceCancel context.CancelFunc
	tasks          tasks.Group
}

func NewDeposit(cfg *config.Config, db *database.DB, btcClient *rpcclient.WalletChainBtcClient, shutdown context.CancelCauseFunc) (*Deposit, error) {

	return nil, nil
}
