package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	btcwallet "github.com/SlimeMutation/btc-wallet-go"
	"github.com/SlimeMutation/btc-wallet-go/common/cliapp"
	"github.com/SlimeMutation/btc-wallet-go/common/opio"
	"github.com/SlimeMutation/btc-wallet-go/config"
	"github.com/SlimeMutation/btc-wallet-go/database"
	flags2 "github.com/SlimeMutation/btc-wallet-go/flags"
	"github.com/SlimeMutation/btc-wallet-go/rpcclient"
	"github.com/SlimeMutation/btc-wallet-go/rpcclient/chain-client/btc"
	"github.com/SlimeMutation/btc-wallet-go/services"
)

const (
	POLLING_INTERVAL     = 1 * time.Second
	MAX_RPC_MESSAGE_SIZE = 1024 * 1024 * 300
)

func runBtcWallet(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Info("exec wallet sync")
	cfg, err := config.LoadConfig(ctx)
	fmt.Println()
	if err != nil {
		log.Error("failed to load config", "err", err)
		return nil, err
	}
	return btcwallet.NewBtcWallet(ctx.Context, &cfg, shutdown)
}

func runRpc(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	fmt.Println("running grpc server...")
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Error("failed to load config", "err", err)
		return nil, err
	}
	grpcServerCfg := &services.BusinessMiddleConfig{
		GrpcHostname: cfg.RpcServer.Host,
		GrpcPort:     cfg.RpcServer.Port,
	}
	db, err := database.NewDB(ctx.Context, cfg.MasterDB)
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
	return services.NewBusinessMiddleWireServices(db, grpcServerCfg, btcClient)
}

func runMigrations(ctx *cli.Context) error {
	ctx.Context = opio.CancelOnInterrupt(ctx.Context)
	log.Info("running migrations...")
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}
	db, err := database.NewDB(ctx.Context, cfg.MasterDB)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return err
	}
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			log.Error("fail to close database", "err", err)
		}
	}(db)
	return db.ExecuteSQLMigration(cfg.Migrations)
}

func NewCli(GitCommit string, GitData string) *cli.App {
	flags := flags2.Flags
	return &cli.App{
		Version:              params.VersionWithCommit(GitCommit, GitData),
		Description:          "An exchange wallet scanner services with rpc and rest api server",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "rpc",
				Flags:       flags,
				Description: "Run rpc services",
				Action:      cliapp.LifecycleCmd(runRpc),
			},
			{
				Name:        "sync",
				Flags:       flags,
				Description: "Run rpc scanner wallet chain node",
				Action:      cliapp.LifecycleCmd(runBtcWallet),
			},
			{
				Name:        "migrate",
				Flags:       flags,
				Description: "Run database migrations",
				Action:      runMigrations,
			},
			{
				Name:        "version",
				Description: "Show project version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}
