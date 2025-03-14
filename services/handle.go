package services

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"

	"github.com/SlimeMutation/btc-wallet-go/database"
	"github.com/SlimeMutation/btc-wallet-go/database/dynamic"
	dal_wallet_go "github.com/SlimeMutation/btc-wallet-go/protobuf/dal-wallet-go"
)

const (
	ChainName = "Bitcoin"
	Network   = "mainnet"
)

func (bws *BusinessMiddleWireServices) BusinessRegister(ctx context.Context, request *dal_wallet_go.BusinessRegisterRequest) (*dal_wallet_go.BusinessRegisterResponse, error) {
	if request.RequestId == "" || request.NotifyUrl == "" {
		return &dal_wallet_go.BusinessRegisterResponse{
			Code: dal_wallet_go.ReturnCode_ERROR,
			Msg:  "invalid params",
		}, nil
	}
	business := &database.Business{
		GUID:        uuid.New(),
		BusinessUid: request.RequestId,
		NotifyUrl:   request.NotifyUrl,
		Timestamp:   uint64(time.Now().Unix()),
	}
	err := bws.db.Business.StoreBusiness(business)
	if err != nil {
		log.Error("store business fail", "err", err)
		return &dal_wallet_go.BusinessRegisterResponse{
			Code: dal_wallet_go.ReturnCode_ERROR,
			Msg:  "store db fail",
		}, nil
	}
	dynamic.CreateTableFromTemplate(request.RequestId, bws.db)
	return &dal_wallet_go.BusinessRegisterResponse{
		Code: dal_wallet_go.ReturnCode_SUCCESS,
		Msg:  "config business success",
	}, nil
}

// ExportAddressesByPublicKeys todo change to tx
// ExportAddressesByPublicKeys
func (bws *BusinessMiddleWireServices) ExportAddressesByPublicKeys(ctx context.Context, request *dal_wallet_go.ExportAddressesRequest) (*dal_wallet_go.ExportAddressesResponse, error) {
	return &dal_wallet_go.ExportAddressesResponse{
		Code:      dal_wallet_go.ReturnCode_SUCCESS,
		Msg:       "generate address success",
		Addresses: nil,
	}, nil
}

func (bws *BusinessMiddleWireServices) BuildUnSignTransaction(ctx context.Context, request *dal_wallet_go.UnSignWithdrawTransactionRequest) (*dal_wallet_go.UnSignWithdrawTransactionResponse, error) {
	return nil, nil
}

func (bws *BusinessMiddleWireServices) BuildSignedTransaction(ctx context.Context, request *dal_wallet_go.SignedWithdrawTransactionRequest) (*dal_wallet_go.SignedWithdrawTransactionResponse, error) {
	return nil, nil
}
