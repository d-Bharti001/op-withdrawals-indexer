package dbmodels

import (
	"fmt"
	"op-withdrawals-indexer/internal/database/dbtypes"
	"op-withdrawals-indexer/internal/database/models"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type WithdrawalStatus string

const (
	WithdrawalInitiated        WithdrawalStatus = "initiated"
	WithdrawalProven           WithdrawalStatus = "proven"
	WithdrawalFinalizedSuccess WithdrawalStatus = "finalized_successful"
	WithdrawalFinalizedFail    WithdrawalStatus = "finalized_failed"
)

type WithdrawalKind string

const (
	CustomWithdrawal  WithdrawalKind = "custom"
	ETHWithdrawal     WithdrawalKind = "eth"
	ERC20Withdrawal   WithdrawalKind = "erc20"
	ERC721Withdrawal  WithdrawalKind = "erc721"
	ERC1155Withdrawal WithdrawalKind = "erc1155"
)

type WithdrawalDetails struct {

	// --- models.Withdrawal ---

	WithdrawalHash common.Hash `json:"withdrawal_hash"`
	ChainID        uint64      `json:"chain_id"`

	Nonce    *string        `json:"withdrawal_nonce"`
	Sender   common.Address `json:"withdrawal_sender"`
	Target   common.Address `json:"withdrawal_target"`
	Value    *string        `json:"withdrawal_value"`
	GasLimit *string        `json:"withdrawal_gas_limit"`
	Data     hexutil.Bytes  `json:"withdrawal_data"`

	TxHash         common.Hash `json:"tx_hash"`
	BlockNumber    uint64      `json:"block_number"`
	BlockTimestamp uint64      `json:"block_timestamp"`

	// --- models.WithdrawalInformation ---

	DecodedAction *string `json:"decoded_action,omitempty"`

	TokenAddress *common.Address `json:"token_address,omitempty"`
	FromAddress  *common.Address `json:"from_address,omitempty"`
	ToAddress    *common.Address `json:"to_address,omitempty"`
	TokenID      *string         `json:"token_id,omitempty"`
	Amount       *string         `json:"amount,omitempty"`

	// --- models.Token ---
	TokenChainID  *uint64            `json:"token_chain_id,omitempty"`
	TokenClass    *models.TokenClass `json:"token_class,omitempty"`
	TokenName     *string            `json:"token_name,omitempty"`
	TokenSymbol   *string            `json:"token_symbol,omitempty"`
	TokenDecimals *uint64            `json:"token_decimals,omitempty"`

	// --- extra fields ---

	Status WithdrawalStatus `json:"status"`
	Kind   WithdrawalKind   `json:"kind"`
}

type WithdrawalDetailsDBRow struct {
	TotalCount uint64

	// --- models.WithdrawalDBRow ---

	WithdrawalHash dbtypes.NullableHash
	ChainID        *uint64

	Nonce    dbtypes.NullableU256
	Sender   dbtypes.NullableAddress
	Target   dbtypes.NullableAddress
	Value    dbtypes.NullableU256
	GasLimit dbtypes.NullableU256
	Data     dbtypes.NullableBytes

	TxHash         dbtypes.NullableHash
	BlockNumber    *uint64
	BlockTimestamp *uint64

	// --- models.WithdrawalInformationDBRow ---

	DecodedAction *string

	TokenAddress dbtypes.NullableAddress
	FromAddress  dbtypes.NullableAddress
	ToAddress    dbtypes.NullableAddress
	TokenID      dbtypes.NullableU256
	Amount       dbtypes.NullableU256

	// --- models.TokenDBRow ---

	TokenChainID  *uint64
	TokenClass    *string
	TokenName     *string
	TokenSymbol   *string
	TokenDecimals *uint64

	// --- extra fields ---

	Proven              *bool
	FinalizationSuccess *bool
}

func (row *WithdrawalDetailsDBRow) ToDomainModel() (*WithdrawalDetails, uint64, error) {

	if row.WithdrawalHash.Common() == nil {
		return nil, row.TotalCount, nil
	}

	result := new(WithdrawalDetails)

	// If withdrawal_hash is not null, some columns are assumed to exist:
	result.WithdrawalHash = *row.WithdrawalHash.Common()
	result.ChainID = *row.ChainID
	result.Nonce = stringPtr(row.Nonce.BigInt().String())
	result.Sender = *row.Sender.Common()
	result.Target = *row.Target.Common()
	result.Value = stringPtr(row.Value.BigInt().String())
	result.GasLimit = stringPtr(row.GasLimit.BigInt().String())
	result.Data = *row.Data.Common()
	result.TxHash = *row.TxHash.Common()
	result.BlockNumber = *row.BlockNumber
	result.BlockTimestamp = *row.BlockTimestamp

	if row.DecodedAction != nil {

		switch *row.DecodedAction {
		case string(models.EthTransferAction):
			result.Kind = ETHWithdrawal
		case string(models.ERC20TransferAction):
			result.Kind = ERC20Withdrawal
		case string(models.ERC721TransferAction):
			result.Kind = ERC721Withdrawal
		case string(models.ERC1155TransferAction):
			result.Kind = ERC1155Withdrawal
		default:
			return nil, 0, fmt.Errorf("unrecognized withdrawal action %s", *row.DecodedAction)
		}

		result.DecodedAction = row.DecodedAction
		result.TokenAddress = row.TokenAddress.Common()
		result.FromAddress = row.FromAddress.Common()
		result.ToAddress = row.ToAddress.Common()
		if row.TokenID.BigInt() != nil {
			result.TokenID = stringPtr(row.TokenID.BigInt().String())
		}
		if row.Amount.BigInt() != nil {
			result.Amount = stringPtr(row.Amount.BigInt().String())
		}

	} else {
		result.Kind = CustomWithdrawal
	}

	if row.TokenChainID != nil && row.TokenClass != nil {
		var tokenClass models.TokenClass

		switch *row.TokenClass {
		case string(models.ERC20Token):
			tokenClass = models.ERC20Token
		case string(models.ERC721Token):
			tokenClass = models.ERC721Token
		case string(models.ERC1155Token):
			tokenClass = models.ERC1155Token
		default:
			return nil, 0, fmt.Errorf("unrecognized token class %s", *row.TokenClass)
		}

		result.TokenChainID = row.TokenChainID
		result.TokenClass = &tokenClass
		result.TokenName = row.TokenName
		result.TokenSymbol = row.TokenSymbol
		result.TokenDecimals = row.TokenDecimals
	}

	if row.FinalizationSuccess != nil {
		if *row.FinalizationSuccess {
			result.Status = WithdrawalFinalizedSuccess
		} else {
			result.Status = WithdrawalFinalizedFail
		}
	} else if row.Proven != nil && *row.Proven == true {
		result.Status = WithdrawalProven
	} else {
		result.Status = WithdrawalInitiated
	}

	return result, row.TotalCount, nil
}

func stringPtr(val string) (ptr *string) {
	return &val
}
