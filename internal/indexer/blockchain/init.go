package blockchain

import (
	"context"
	"errors"

	"op-withdrawals-indexer/internal/chainclient"
	"op-withdrawals-indexer/internal/contracts/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

func NewL1Chain(ctx context.Context, cfg L1ChainInitConfig) (*L1Chain, error) {
	ctx, cancel := context.WithTimeout(ctx, DefaultRpcTimeout)
	defer cancel()

	l1Client, l1ChainId, err := chainclient.DialChain(chainclient.DialChainParams{
		Ctx:     ctx,
		URL:     cfg.RPCUrl,
		ChainID: cfg.ID,
		Name:    cfg.Name,
	})
	if err != nil {
		return nil, err
	}

	l1Chain := L1Chain{
		Chain: Chain{
			id:                     l1ChainId,
			name:                   cfg.Name,
			rpcClient:              l1Client,
			blockConfirmationDepth: cfg.BlockConfirmationDepth,
		},
		iOptimismPortal2: bindings.NewIOptimismPortal2(),
	}

	return &l1Chain, nil
}

func NewL2Chain(ctx context.Context, cfg L2ChainInitConfig) (*L2Chain, error) {
	ctx, cancel := context.WithTimeout(ctx, DefaultRpcTimeout)
	defer cancel()

	l2Client, l2ChainId, err := chainclient.DialChain(chainclient.DialChainParams{
		Ctx:     ctx,
		URL:     cfg.RPCUrl,
		ChainID: cfg.ID,
		Name:    cfg.Name,
	})
	if err != nil {
		return nil, err
	}

	l1Client, l1ChainId, err := chainclient.DialChain(chainclient.DialChainParams{
		Ctx:     ctx,
		URL:     cfg.L1RPCUrl,
		ChainID: cfg.L1ChainID,
		Name:    "L1 chain for " + cfg.Name,
	})
	if err != nil {
		return nil, err
	}

	sysConfigContract := bindings.NewISystemConfig()
	sysConfigInstance := sysConfigContract.Instance(l1Client, common.HexToAddress(cfg.SystemConfigAddr))

	callOpts := &bind.CallOpts{}

	sysConfigL2ChainId, err := bind.Call(
		sysConfigInstance,
		callOpts,
		sysConfigContract.PackL2ChainId(),
		sysConfigContract.UnpackL2ChainId,
	)
	if err != nil {
		return nil, err
	}
	if sysConfigL2ChainId.Uint64() != cfg.ID {
		return nil, errors.New("system config l2 chain id mismatch")
	}

	startBlock, err := bind.Call(
		sysConfigInstance,
		callOpts,
		sysConfigContract.PackStartBlock(),
		sysConfigContract.UnpackStartBlock,
	)
	if err != nil {
		return nil, err
	}

	optimismPortalAddr, err := bind.Call(
		sysConfigInstance,
		callOpts,
		sysConfigContract.PackOptimismPortal(),
		sysConfigContract.UnpackOptimismPortal,
	)
	if err != nil {
		return nil, err
	}

	l1CrossDomainMessengerAddr, err := bind.Call(
		sysConfigInstance,
		callOpts,
		sysConfigContract.PackL1CrossDomainMessenger(),
		sysConfigContract.UnpackL1CrossDomainMessenger,
	)
	if err != nil {
		return nil, err
	}

	l1StandardBridgeAddr, err := bind.Call(
		sysConfigInstance,
		callOpts,
		sysConfigContract.PackL1StandardBridge(),
		sysConfigContract.UnpackL1StandardBridge,
	)
	if err != nil {
		return nil, err
	}

	l1ERC721BridgeAddr, err := bind.Call(
		sysConfigInstance,
		callOpts,
		sysConfigContract.PackL1ERC721Bridge(),
		sysConfigContract.UnpackL1ERC721Bridge,
	)
	if err != nil {
		return nil, err
	}

	l2Chain := L2Chain{
		Chain: Chain{
			id:                     l2ChainId,
			name:                   cfg.Name,
			rpcClient:              l2Client,
			blockConfirmationDepth: cfg.BlockConfirmationDepth,
		},
		LinkedL1Details: LinkedL1Details{
			l1ChainId:          l1ChainId,
			l1StartBlockNumber: startBlock.Uint64(),
			L1ContractAddresses: L1ContractAddresses{
				optimismPortalAddr:         optimismPortalAddr,
				l1CrossDomainMessengerAddr: l1CrossDomainMessengerAddr,
				l1StandardBridgeAddr:       l1StandardBridgeAddr,
				l1ERC721BridgeAddr:         l1ERC721BridgeAddr,
			},
		},
		iL2ToL1MessagePasser: bindings.NewIL2ToL1MessagePasser(),
	}

	return &l2Chain, nil
}
