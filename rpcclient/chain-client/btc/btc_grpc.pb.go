// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/btc.proto

package btc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WalletBtcService_ConvertAddress_FullMethodName          = "/WalletBtcService/convertAddress"
	WalletBtcService_ValidAddress_FullMethodName            = "/WalletBtcService/validAddress"
	WalletBtcService_GetFee_FullMethodName                  = "/WalletBtcService/getFee"
	WalletBtcService_GetAccount_FullMethodName              = "/WalletBtcService/getAccount"
	WalletBtcService_GetUnspentOutputs_FullMethodName       = "/WalletBtcService/getUnspentOutputs"
	WalletBtcService_GetBlockByNumber_FullMethodName        = "/WalletBtcService/getBlockByNumber"
	WalletBtcService_GetBlockByHash_FullMethodName          = "/WalletBtcService/getBlockByHash"
	WalletBtcService_GetBlockHeaderByHash_FullMethodName    = "/WalletBtcService/getBlockHeaderByHash"
	WalletBtcService_GetBlockHeaderByNumber_FullMethodName  = "/WalletBtcService/getBlockHeaderByNumber"
	WalletBtcService_SendTx_FullMethodName                  = "/WalletBtcService/SendTx"
	WalletBtcService_GetTxByAddress_FullMethodName          = "/WalletBtcService/getTxByAddress"
	WalletBtcService_GetTxByHash_FullMethodName             = "/WalletBtcService/getTxByHash"
	WalletBtcService_BuildUnSignTransaction_FullMethodName  = "/WalletBtcService/buildUnSignTransaction"
	WalletBtcService_BuildSignedTransaction_FullMethodName  = "/WalletBtcService/buildSignedTransaction"
	WalletBtcService_DecodeTransaction_FullMethodName       = "/WalletBtcService/decodeTransaction"
	WalletBtcService_VerifySignedTransaction_FullMethodName = "/WalletBtcService/verifySignedTransaction"
)

// WalletBtcServiceClient is the client API for WalletBtcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletBtcServiceClient interface {
	ConvertAddress(ctx context.Context, in *ConvertAddressRequest, opts ...grpc.CallOption) (*ConvertAddressResponse, error)
	ValidAddress(ctx context.Context, in *ValidAddressRequest, opts ...grpc.CallOption) (*ValidAddressResponse, error)
	GetFee(ctx context.Context, in *FeeRequest, opts ...grpc.CallOption) (*FeeResponse, error)
	GetAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	GetUnspentOutputs(ctx context.Context, in *UnspentOutputsRequest, opts ...grpc.CallOption) (*UnspentOutputsResponse, error)
	GetBlockByNumber(ctx context.Context, in *BlockNumberRequest, opts ...grpc.CallOption) (*BlockResponse, error)
	GetBlockByHash(ctx context.Context, in *BlockHashRequest, opts ...grpc.CallOption) (*BlockResponse, error)
	GetBlockHeaderByHash(ctx context.Context, in *BlockHeaderHashRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error)
	GetBlockHeaderByNumber(ctx context.Context, in *BlockHeaderNumberRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error)
	SendTx(ctx context.Context, in *SendTxRequest, opts ...grpc.CallOption) (*SendTxResponse, error)
	GetTxByAddress(ctx context.Context, in *TxAddressRequest, opts ...grpc.CallOption) (*TxAddressResponse, error)
	GetTxByHash(ctx context.Context, in *TxHashRequest, opts ...grpc.CallOption) (*TxHashResponse, error)
	BuildUnSignTransaction(ctx context.Context, in *UnSignTransactionRequest, opts ...grpc.CallOption) (*UnSignTransactionResponse, error)
	BuildSignedTransaction(ctx context.Context, in *SignedTransactionRequest, opts ...grpc.CallOption) (*SignedTransactionResponse, error)
	DecodeTransaction(ctx context.Context, in *DecodeTransactionRequest, opts ...grpc.CallOption) (*DecodeTransactionResponse, error)
	VerifySignedTransaction(ctx context.Context, in *VerifyTransactionRequest, opts ...grpc.CallOption) (*VerifyTransactionResponse, error)
}

type walletBtcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletBtcServiceClient(cc grpc.ClientConnInterface) WalletBtcServiceClient {
	return &walletBtcServiceClient{cc}
}

func (c *walletBtcServiceClient) ConvertAddress(ctx context.Context, in *ConvertAddressRequest, opts ...grpc.CallOption) (*ConvertAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConvertAddressResponse)
	err := c.cc.Invoke(ctx, WalletBtcService_ConvertAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletBtcServiceClient) ValidAddress(ctx context.Context, in *ValidAddressRequest, opts ...grpc.CallOption) (*ValidAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidAddressResponse)
	err := c.cc.Invoke(ctx, WalletBtcService_ValidAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletBtcServiceClient) GetFee(ctx context.Context, in *FeeRequest, opts ...grpc.CallOption) (*FeeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FeeResponse)
	err := c.cc.Invoke(ctx, WalletBtcService_GetFee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletBtcServiceClient) GetAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, WalletBtcService_GetAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletBtcServiceClient) GetUnspentOutputs(ctx context.Context, in *UnspentOutputsRequest, opts ...grpc.CallOption) (*UnspentOutputsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnspentOutputsResponse)
	err := c.cc.Invoke(ctx, WalletBtcService_GetUnspentOutputs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletBtcServiceClient) GetBlockByNumber(ctx context.Context, in *BlockNumberRequest, opts ...grpc.CallOption) (*BlockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, WalletBtcService_GetBlockByNumber_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletBtcServiceClient) GetBlockByHash(ctx context.Context, in *BlockHashRequest, opts ...grpc.CallOption) (*BlockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, WalletBtcService_GetBlockByHash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletBtcServiceClient) GetBlockHeaderByHash(ctx context.Context, in *BlockHeaderHashRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockHeaderResponse)
	err := c.cc.Invoke(ctx, WalletBtcService_GetBlockHeaderByHash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletBtcServiceClient) GetBlockHeaderByNumber(ctx context.Context, in *BlockHeaderNumberRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockHeaderResponse)
	err := c.cc.Invoke(ctx, WalletBtcService_GetBlockHeaderByNumber_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletBtcServiceClient) SendTx(ctx context.Context, in *SendTxRequest, opts ...grpc.CallOption) (*SendTxResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendTxResponse)
	err := c.cc.Invoke(ctx, WalletBtcService_SendTx_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletBtcServiceClient) GetTxByAddress(ctx context.Context, in *TxAddressRequest, opts ...grpc.CallOption) (*TxAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TxAddressResponse)
	err := c.cc.Invoke(ctx, WalletBtcService_GetTxByAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletBtcServiceClient) GetTxByHash(ctx context.Context, in *TxHashRequest, opts ...grpc.CallOption) (*TxHashResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TxHashResponse)
	err := c.cc.Invoke(ctx, WalletBtcService_GetTxByHash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletBtcServiceClient) BuildUnSignTransaction(ctx context.Context, in *UnSignTransactionRequest, opts ...grpc.CallOption) (*UnSignTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnSignTransactionResponse)
	err := c.cc.Invoke(ctx, WalletBtcService_BuildUnSignTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletBtcServiceClient) BuildSignedTransaction(ctx context.Context, in *SignedTransactionRequest, opts ...grpc.CallOption) (*SignedTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignedTransactionResponse)
	err := c.cc.Invoke(ctx, WalletBtcService_BuildSignedTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletBtcServiceClient) DecodeTransaction(ctx context.Context, in *DecodeTransactionRequest, opts ...grpc.CallOption) (*DecodeTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DecodeTransactionResponse)
	err := c.cc.Invoke(ctx, WalletBtcService_DecodeTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletBtcServiceClient) VerifySignedTransaction(ctx context.Context, in *VerifyTransactionRequest, opts ...grpc.CallOption) (*VerifyTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyTransactionResponse)
	err := c.cc.Invoke(ctx, WalletBtcService_VerifySignedTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBtcServiceServer is the server API for WalletBtcService service.
// All implementations should embed UnimplementedWalletBtcServiceServer
// for forward compatibility.
type WalletBtcServiceServer interface {
	ConvertAddress(context.Context, *ConvertAddressRequest) (*ConvertAddressResponse, error)
	ValidAddress(context.Context, *ValidAddressRequest) (*ValidAddressResponse, error)
	GetFee(context.Context, *FeeRequest) (*FeeResponse, error)
	GetAccount(context.Context, *AccountRequest) (*AccountResponse, error)
	GetUnspentOutputs(context.Context, *UnspentOutputsRequest) (*UnspentOutputsResponse, error)
	GetBlockByNumber(context.Context, *BlockNumberRequest) (*BlockResponse, error)
	GetBlockByHash(context.Context, *BlockHashRequest) (*BlockResponse, error)
	GetBlockHeaderByHash(context.Context, *BlockHeaderHashRequest) (*BlockHeaderResponse, error)
	GetBlockHeaderByNumber(context.Context, *BlockHeaderNumberRequest) (*BlockHeaderResponse, error)
	SendTx(context.Context, *SendTxRequest) (*SendTxResponse, error)
	GetTxByAddress(context.Context, *TxAddressRequest) (*TxAddressResponse, error)
	GetTxByHash(context.Context, *TxHashRequest) (*TxHashResponse, error)
	BuildUnSignTransaction(context.Context, *UnSignTransactionRequest) (*UnSignTransactionResponse, error)
	BuildSignedTransaction(context.Context, *SignedTransactionRequest) (*SignedTransactionResponse, error)
	DecodeTransaction(context.Context, *DecodeTransactionRequest) (*DecodeTransactionResponse, error)
	VerifySignedTransaction(context.Context, *VerifyTransactionRequest) (*VerifyTransactionResponse, error)
}

// UnimplementedWalletBtcServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWalletBtcServiceServer struct{}

func (UnimplementedWalletBtcServiceServer) ConvertAddress(context.Context, *ConvertAddressRequest) (*ConvertAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertAddress not implemented")
}
func (UnimplementedWalletBtcServiceServer) ValidAddress(context.Context, *ValidAddressRequest) (*ValidAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidAddress not implemented")
}
func (UnimplementedWalletBtcServiceServer) GetFee(context.Context, *FeeRequest) (*FeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFee not implemented")
}
func (UnimplementedWalletBtcServiceServer) GetAccount(context.Context, *AccountRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedWalletBtcServiceServer) GetUnspentOutputs(context.Context, *UnspentOutputsRequest) (*UnspentOutputsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnspentOutputs not implemented")
}
func (UnimplementedWalletBtcServiceServer) GetBlockByNumber(context.Context, *BlockNumberRequest) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockByNumber not implemented")
}
func (UnimplementedWalletBtcServiceServer) GetBlockByHash(context.Context, *BlockHashRequest) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockByHash not implemented")
}
func (UnimplementedWalletBtcServiceServer) GetBlockHeaderByHash(context.Context, *BlockHeaderHashRequest) (*BlockHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeaderByHash not implemented")
}
func (UnimplementedWalletBtcServiceServer) GetBlockHeaderByNumber(context.Context, *BlockHeaderNumberRequest) (*BlockHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeaderByNumber not implemented")
}
func (UnimplementedWalletBtcServiceServer) SendTx(context.Context, *SendTxRequest) (*SendTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTx not implemented")
}
func (UnimplementedWalletBtcServiceServer) GetTxByAddress(context.Context, *TxAddressRequest) (*TxAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxByAddress not implemented")
}
func (UnimplementedWalletBtcServiceServer) GetTxByHash(context.Context, *TxHashRequest) (*TxHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxByHash not implemented")
}
func (UnimplementedWalletBtcServiceServer) BuildUnSignTransaction(context.Context, *UnSignTransactionRequest) (*UnSignTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildUnSignTransaction not implemented")
}
func (UnimplementedWalletBtcServiceServer) BuildSignedTransaction(context.Context, *SignedTransactionRequest) (*SignedTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildSignedTransaction not implemented")
}
func (UnimplementedWalletBtcServiceServer) DecodeTransaction(context.Context, *DecodeTransactionRequest) (*DecodeTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecodeTransaction not implemented")
}
func (UnimplementedWalletBtcServiceServer) VerifySignedTransaction(context.Context, *VerifyTransactionRequest) (*VerifyTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifySignedTransaction not implemented")
}
func (UnimplementedWalletBtcServiceServer) testEmbeddedByValue() {}

// UnsafeWalletBtcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletBtcServiceServer will
// result in compilation errors.
type UnsafeWalletBtcServiceServer interface {
	mustEmbedUnimplementedWalletBtcServiceServer()
}

func RegisterWalletBtcServiceServer(s grpc.ServiceRegistrar, srv WalletBtcServiceServer) {
	// If the following call pancis, it indicates UnimplementedWalletBtcServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WalletBtcService_ServiceDesc, srv)
}

func _WalletBtcService_ConvertAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletBtcServiceServer).ConvertAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletBtcService_ConvertAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletBtcServiceServer).ConvertAddress(ctx, req.(*ConvertAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletBtcService_ValidAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletBtcServiceServer).ValidAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletBtcService_ValidAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletBtcServiceServer).ValidAddress(ctx, req.(*ValidAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletBtcService_GetFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletBtcServiceServer).GetFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletBtcService_GetFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletBtcServiceServer).GetFee(ctx, req.(*FeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletBtcService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletBtcServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletBtcService_GetAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletBtcServiceServer).GetAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletBtcService_GetUnspentOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnspentOutputsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletBtcServiceServer).GetUnspentOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletBtcService_GetUnspentOutputs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletBtcServiceServer).GetUnspentOutputs(ctx, req.(*UnspentOutputsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletBtcService_GetBlockByNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletBtcServiceServer).GetBlockByNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletBtcService_GetBlockByNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletBtcServiceServer).GetBlockByNumber(ctx, req.(*BlockNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletBtcService_GetBlockByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletBtcServiceServer).GetBlockByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletBtcService_GetBlockByHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletBtcServiceServer).GetBlockByHash(ctx, req.(*BlockHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletBtcService_GetBlockHeaderByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHeaderHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletBtcServiceServer).GetBlockHeaderByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletBtcService_GetBlockHeaderByHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletBtcServiceServer).GetBlockHeaderByHash(ctx, req.(*BlockHeaderHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletBtcService_GetBlockHeaderByNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHeaderNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletBtcServiceServer).GetBlockHeaderByNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletBtcService_GetBlockHeaderByNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletBtcServiceServer).GetBlockHeaderByNumber(ctx, req.(*BlockHeaderNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletBtcService_SendTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletBtcServiceServer).SendTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletBtcService_SendTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletBtcServiceServer).SendTx(ctx, req.(*SendTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletBtcService_GetTxByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletBtcServiceServer).GetTxByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletBtcService_GetTxByAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletBtcServiceServer).GetTxByAddress(ctx, req.(*TxAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletBtcService_GetTxByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletBtcServiceServer).GetTxByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletBtcService_GetTxByHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletBtcServiceServer).GetTxByHash(ctx, req.(*TxHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletBtcService_BuildUnSignTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnSignTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletBtcServiceServer).BuildUnSignTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletBtcService_BuildUnSignTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletBtcServiceServer).BuildUnSignTransaction(ctx, req.(*UnSignTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletBtcService_BuildSignedTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletBtcServiceServer).BuildSignedTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletBtcService_BuildSignedTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletBtcServiceServer).BuildSignedTransaction(ctx, req.(*SignedTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletBtcService_DecodeTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletBtcServiceServer).DecodeTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletBtcService_DecodeTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletBtcServiceServer).DecodeTransaction(ctx, req.(*DecodeTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletBtcService_VerifySignedTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletBtcServiceServer).VerifySignedTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletBtcService_VerifySignedTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletBtcServiceServer).VerifySignedTransaction(ctx, req.(*VerifyTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletBtcService_ServiceDesc is the grpc.ServiceDesc for WalletBtcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletBtcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WalletBtcService",
	HandlerType: (*WalletBtcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "convertAddress",
			Handler:    _WalletBtcService_ConvertAddress_Handler,
		},
		{
			MethodName: "validAddress",
			Handler:    _WalletBtcService_ValidAddress_Handler,
		},
		{
			MethodName: "getFee",
			Handler:    _WalletBtcService_GetFee_Handler,
		},
		{
			MethodName: "getAccount",
			Handler:    _WalletBtcService_GetAccount_Handler,
		},
		{
			MethodName: "getUnspentOutputs",
			Handler:    _WalletBtcService_GetUnspentOutputs_Handler,
		},
		{
			MethodName: "getBlockByNumber",
			Handler:    _WalletBtcService_GetBlockByNumber_Handler,
		},
		{
			MethodName: "getBlockByHash",
			Handler:    _WalletBtcService_GetBlockByHash_Handler,
		},
		{
			MethodName: "getBlockHeaderByHash",
			Handler:    _WalletBtcService_GetBlockHeaderByHash_Handler,
		},
		{
			MethodName: "getBlockHeaderByNumber",
			Handler:    _WalletBtcService_GetBlockHeaderByNumber_Handler,
		},
		{
			MethodName: "SendTx",
			Handler:    _WalletBtcService_SendTx_Handler,
		},
		{
			MethodName: "getTxByAddress",
			Handler:    _WalletBtcService_GetTxByAddress_Handler,
		},
		{
			MethodName: "getTxByHash",
			Handler:    _WalletBtcService_GetTxByHash_Handler,
		},
		{
			MethodName: "buildUnSignTransaction",
			Handler:    _WalletBtcService_BuildUnSignTransaction_Handler,
		},
		{
			MethodName: "buildSignedTransaction",
			Handler:    _WalletBtcService_BuildSignedTransaction_Handler,
		},
		{
			MethodName: "decodeTransaction",
			Handler:    _WalletBtcService_DecodeTransaction_Handler,
		},
		{
			MethodName: "verifySignedTransaction",
			Handler:    _WalletBtcService_VerifySignedTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/btc.proto",
}
