package blockchain

import (
	"context"
	"errors"

	"op-withdrawals-indexer/internal/contracts/bindings"
	"op-withdrawals-indexer/internal/contracts/predeploys"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func InitNewChains(cfg ChainsInitConfig) (*L2Chain, *L1Chain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultRpcTimeout)
	defer cancel()

	l1Client, err := ethclient.DialContext(ctx, cfg.L1RPC)
	if err != nil {
		return nil, nil, err
	}
	l1ChainId, err := l1Client.ChainID(ctx)
	if err != nil {
		return nil, nil, err
	}
	if l1ChainId.Uint64() != cfg.L1ChainID {
		return nil, nil, errors.New("l1 chain id different than expected")
	}

	l2Client, err := ethclient.DialContext(ctx, cfg.L2RPC)
	if err != nil {
		return nil, nil, err
	}
	l2ChainId, err := l2Client.ChainID(ctx)
	if err != nil {
		return nil, nil, err
	}
	if l2ChainId.Uint64() != cfg.L2ChainID {
		return nil, nil, errors.New("l2 chain id different than expected")
	}

	sysConfigContract := bindings.NewISystemConfig()
	sysConfigInstance := sysConfigContract.Instance(l1Client, common.HexToAddress(cfg.SystemConfigAddr))

	callOpts := &bind.CallOpts{}

	// Verify chain id with the SystemConfig contract
	sysConfigL2ChainId, err := bind.Call(
		sysConfigInstance,
		callOpts,
		sysConfigContract.PackL2ChainId(),
		sysConfigContract.UnpackL2ChainId,
	)
	if err != nil {
		return nil, nil, err
	}
	if l2ChainId.Cmp(sysConfigL2ChainId) != 0 {
		return nil, nil, errors.New("system config l2 chain id mismatch")
	}

	startBlock, err := bind.Call(
		sysConfigInstance,
		callOpts,
		sysConfigContract.PackStartBlock(),
		sysConfigContract.UnpackStartBlock,
	)
	if err != nil {
		return nil, nil, err
	}

	optimismPortalAddr, err := bind.Call(
		sysConfigInstance,
		callOpts,
		sysConfigContract.PackOptimismPortal(),
		sysConfigContract.UnpackOptimismPortal,
	)
	if err != nil {
		return nil, nil, err
	}

	l1CrossDomainMessengerAddr, err := bind.Call(
		sysConfigInstance,
		callOpts,
		sysConfigContract.PackL1CrossDomainMessenger(),
		sysConfigContract.UnpackL1CrossDomainMessenger,
	)
	if err != nil {
		return nil, nil, err
	}

	l1StandardBridgeAddr, err := bind.Call(
		sysConfigInstance,
		callOpts,
		sysConfigContract.PackL1StandardBridge(),
		sysConfigContract.UnpackL1StandardBridge,
	)
	if err != nil {
		return nil, nil, err
	}

	l1ERC721BridgeAddr, err := bind.Call(
		sysConfigInstance,
		callOpts,
		sysConfigContract.PackL1ERC721Bridge(),
		sysConfigContract.UnpackL1ERC721Bridge,
	)
	if err != nil {
		return nil, nil, err
	}

	var l1Contracts L1Contracts

	l1Contracts.OptimismPortal = bindings.NewIOptimismPortal2().Instance(l1Client, optimismPortalAddr)
	l1Contracts.L1CrossDomainMessenger = bindings.NewIOptimismPortal2().Instance(l1Client, l1CrossDomainMessengerAddr)
	l1Contracts.L1StandardBridge = bindings.NewIOptimismPortal2().Instance(l1Client, l1StandardBridgeAddr)
	l1Contracts.L1ERC721Bridge = bindings.NewIOptimismPortal2().Instance(l1Client, l1ERC721BridgeAddr)

	var l2Contracts L2Contracts

	l2Contracts.L2ToL1MessagePasser = bindings.NewIL2ToL1MessagePasser().Instance(l2Client, predeploys.L2ToL1MessagePasserAddr)
	l2Contracts.L2CrossDomainMessenger = bindings.NewIL2CrossDomainMessenger().Instance(l2Client, predeploys.L2CrossDomainMessengerAddr)
	l2Contracts.L2StandardBridge = bindings.NewIL2StandardBridge().Instance(l2Client, predeploys.L2StandardBridgeAddr)
	l2Contracts.L2ERC721Bridge = bindings.NewIL2ERC721Bridge().Instance(l2Client, predeploys.L2ERC721BridgeAddr)

	if cfg.L1BlockConfirmationDepth == nil {
		cfg.L1BlockConfirmationDepth = &DefaultL1ChainBlockConfDepth
	}
	if cfg.L2BlockConfirmationDepth == nil {
		cfg.L2BlockConfirmationDepth = &DefaultL2ChainBlockConfDepth
	}

	l2Chain := L2Chain{
		Chain: Chain{
			ID:                     l2ChainId.Uint64(),
			Name:                   cfg.L2ChainName,
			RPCClient:              l2Client,
			BlockConfirmationDepth: *cfg.L2BlockConfirmationDepth,
			PollingInterval:        DefaultL2PollingInterval,
		},
		L1ChainID:   cfg.L1ChainID,
		L2Contracts: l2Contracts,
	}

	l2ChainOnL1 := L2ChainOnL1{
		L2Chain:     l2Chain,
		StartBlock:  startBlock.Uint64(),
		L1Contracts: l1Contracts,
	}

	l1Chain := L1Chain{
		Chain: Chain{
			ID:                     l1ChainId.Uint64(),
			Name:                   cfg.L1ChainName,
			RPCClient:              l1Client,
			BlockConfirmationDepth: *cfg.L1BlockConfirmationDepth,
			PollingInterval:        DefaultL1PollingInterval,
		},
		L2Chains: []L2ChainOnL1{l2ChainOnL1},
	}

	return &l2Chain, &l1Chain, nil
}
