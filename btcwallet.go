package btc_wallet

import (
	"context"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/SlimeMutation/btc-wallet-go/config"
	"github.com/SlimeMutation/btc-wallet-go/database"
	"github.com/SlimeMutation/btc-wallet-go/rpcclient"
	"github.com/SlimeMutation/btc-wallet-go/rpcclient/chain-client/btc"
	"github.com/SlimeMutation/btc-wallet-go/worker"
)

type BtcWallet struct {
	deposit *worker.Deposit

	shutdown context.CancelCauseFunc
	stopped  atomic.Bool
}

func NewBtcWallet(ctx context.Context, cfg *config.Config, shutdown context.CancelCauseFunc) (*BtcWallet, error) {
	db, err := database.NewDB(ctx, cfg.MasterDB)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return nil, err
	}

	log.Info("Chain btc rpc", "rpc uri", cfg.ChainBtcRpc)
	conn, err := grpc.NewClient(cfg.ChainBtcRpc, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("Connect to da retriever fail", "err", err)
		return nil, err
	}
	client := btc.NewWalletBtcServiceClient(conn)
	btcClient, err := rpcclient.NewWalletChainBtcClient(context.Background(), client, "mainnet")
	if err != nil {
		log.Error("new wallet btc client fail", "err", err)
		return nil, err
	}

	deposit, _ := worker.NewDeposit(cfg, db, btcClient, shutdown)

	out := &BtcWallet{
		deposit:  deposit,
		shutdown: shutdown,
	}

	return out, nil
}

func (ew *BtcWallet) Start(ctx context.Context) error {
	err := ew.deposit.Start()
	if err != nil {
		return err
	}
	return nil
}

func (ew *BtcWallet) Stop(ctx context.Context) error {
	err := ew.deposit.Close()
	if err != nil {
		return err
	}
	return nil
}

func (ew *BtcWallet) Stopped() bool {
	return ew.stopped.Load()
}
