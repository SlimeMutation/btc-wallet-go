package rpcclient

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/log"

	"github.com/SlimeMutation/btc-wallet-go/rpcclient/chain-client/btc"
)

type WalletChainBtcClient struct {
	Ctx               context.Context
	Network           string
	ChainBtcRpcClient btc.WalletBtcServiceClient
}

func NewWalletChainBtcClient(ctx context.Context, rpc btc.WalletBtcServiceClient, network string) (*WalletChainBtcClient, error) {
	log.Info("New chain btc rpc client", "network", network)
	return &WalletChainBtcClient{Ctx: ctx, ChainBtcRpcClient: rpc, Network: network}, nil
}

func (wac *WalletChainBtcClient) ExportAddressByPubKey(format, publicKey string) string {
	req := &btc.ConvertAddressRequest{
		Network:   wac.Network,
		Format:    format,
		PublicKey: publicKey,
	}
	address, err := wac.ChainBtcRpcClient.ConvertAddress(wac.Ctx, req)
	if err != nil {
		log.Error("covert address fail", "err", err)
		return ""
	}
	if address.Code == btc.ReturnCode_ERROR {
		log.Error("covert address fail", "err", err)
		return ""
	}
	return address.Address
}

func (wac *WalletChainBtcClient) GetBlockHeader(height *big.Int) (*BlockHeader, error) {
	req := &btc.BlockHeaderNumberRequest{
		Network: wac.Network,
		Height:  height.Int64(),
	}
	response, err := wac.ChainBtcRpcClient.GetBlockHeaderByNumber(wac.Ctx, req)
	if err != nil {
		log.Error("GetBlockHeader fail", "err", err)
		return nil, err
	}
	if response.Code == btc.ReturnCode_ERROR {
		log.Error("GetBlockHeader fail", "err", err)
		return nil, errors.New("GetBlockHeader fail")
	}
	number, _ := new(big.Int).SetString(response.Number, 10)
	return &BlockHeader{
		Hash:       response.BlockHash,
		ParentHash: response.ParentHash,
		Number:     number,
		Timestamp:  uint64(response.Time),
	}, nil
}
