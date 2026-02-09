package dbmodels

import (
	"fmt"
	"math/big"
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

	Sender common.Address `json:"withdrawal_sender"`
	Target common.Address `json:"withdrawal_target"`
	Data   hexutil.Bytes  `json:"withdrawal_data"`
	Value  *big.Int       `json:"withdrawal_value"`

	TxHash         common.Hash `json:"tx_hash"`
	BlockNumber    uint64      `json:"block_number"`
	BlockTimestamp uint64      `json:"block_timestamp"`

	// --- models.WithdrawalInformation ---

	DecodedAction *string `json:"decoded_action,omitempty"`

	TokenAddress *common.Address `json:"token_address,omitempty"`
	FromAddress  *common.Address `json:"from_address,omitempty"`
	ToAddress    *common.Address `json:"to_address,omitempty"`
	TokenID      *big.Int        `json:"token_id,omitempty"`
	Amount       *big.Int        `json:"amount,omitempty"`

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

	// --- models.WithdrawalDBRow ---

	WithdrawalHash dbtypes.Hash
	ChainID        uint64

	Sender dbtypes.Address
	Target dbtypes.Address
	Data   dbtypes.Bytes
	Value  dbtypes.U256

	TxHash         dbtypes.Hash
	BlockNumber    uint64
	BlockTimestamp uint64

	// --- models.WithdrawalInformationDBRow ---

	DecodedAction dbtypes.NullableString

	TokenAddress dbtypes.NullableAddress
	FromAddress  dbtypes.NullableAddress
	ToAddress    dbtypes.NullableAddress
	TokenID      dbtypes.NullableU256
	Amount       dbtypes.NullableU256

	// --- models.TokenDBRow ---

	TokenChainID  dbtypes.NullableUint64
	TokenClass    dbtypes.NullableString
	TokenName     dbtypes.NullableString
	TokenSymbol   dbtypes.NullableString
	TokenDecimals dbtypes.NullableUint64

	// --- extra fields ---

	Proven              bool
	FinalizationSuccess *bool
}

func (row *WithdrawalDetailsDBRow) ToDomainModel() (*WithdrawalDetails, error) {
	result := new(WithdrawalDetails)

	result.WithdrawalHash = row.WithdrawalHash.Common()
	result.ChainID = row.ChainID
	result.Sender = row.Sender.Common()
	result.Target = row.Target.Common()
	result.Data = row.Data.HexUtilBytes()
	result.Value = row.Value.BigInt()
	result.TxHash = row.TxHash.Common()
	result.BlockNumber = row.BlockNumber
	result.BlockTimestamp = row.BlockTimestamp

	if row.DecodedAction.Val != nil {

		switch *row.DecodedAction.Val {
		case string(models.EthTransferAction):
			result.Kind = ETHWithdrawal
		case string(models.ERC20TransferAction):
			result.Kind = ERC20Withdrawal
		case string(models.ERC721TransferAction):
			result.Kind = ERC721Withdrawal
		case string(models.ERC1155TransferAction):
			result.Kind = ERC1155Withdrawal
		default:
			return nil, fmt.Errorf("unrecognized withdrawal action %s", *row.DecodedAction.Val)
		}

		result.DecodedAction = row.DecodedAction.Val
		result.TokenAddress = row.TokenAddress.Val
		result.FromAddress = row.FromAddress.Common()
		result.ToAddress = row.ToAddress.Common()
		result.TokenID = row.TokenID.BigInt()
		result.Amount = row.Amount.BigInt()

	} else {
		result.Kind = CustomWithdrawal
	}

	if row.TokenChainID.Val != nil && row.TokenClass.Val != nil {
		var tokenClass models.TokenClass

		switch *row.TokenClass.Val {
		case string(models.ERC20Token):
			tokenClass = models.ERC20Token
		case string(models.ERC721Token):
			tokenClass = models.ERC721Token
		case string(models.ERC1155Token):
			tokenClass = models.ERC1155Token
		default:
			return nil, fmt.Errorf("unrecognized token class %s", *row.TokenClass.Val)
		}

		result.TokenChainID = row.TokenChainID.Val
		result.TokenClass = &tokenClass
		result.TokenName = row.TokenName.Val
		result.TokenSymbol = row.TokenSymbol.Val
		result.TokenDecimals = row.TokenDecimals.Val
	}

	if row.FinalizationSuccess != nil {
		if *row.FinalizationSuccess {
			result.Status = WithdrawalFinalizedSuccess
		} else {
			result.Status = WithdrawalFinalizedFail
		}
	} else if row.Proven {
		result.Status = WithdrawalProven
	} else {
		result.Status = WithdrawalInitiated
	}

	return result, nil
}
