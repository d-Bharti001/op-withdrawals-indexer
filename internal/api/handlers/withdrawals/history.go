package withdrawals

import (
	"net/http"
	"strconv"
	"time"

	"op-withdrawals-indexer/internal/api/dbmodels"
	"op-withdrawals-indexer/internal/api/handlers/response"

	"github.com/ethereum/go-ethereum/common"
)

type EnvelopedWithdrawalHistoryResponse struct {
	Message string `json:"message"`
	Result  struct {
		Address     string `json:"address"`
		ChainID     uint64 `json:"source_chain_id"`
		Withdrawals []struct {
			WithdrawalHash string  `json:"withdrawal_hash"`
			ChainID        uint64  `json:"chain_id"`
			Sender         string  `json:"withdrawal_sender"`
			Target         string  `json:"withdrawal_target"`
			Data           string  `json:"withdrawal_data"`
			Value          string  `json:"withdrawal_value"`
			TxHash         string  `json:"tx_hash"`
			BlockNumber    uint64  `json:"block_number"`
			BlockTimestamp uint64  `json:"block_timestamp"`
			DecodedAction  *string `json:"decoded_action,omitempty"`
			TokenAddress   *string `json:"token_address,omitempty"`
			FromAddress    *string `json:"from_address,omitempty"`
			ToAddress      *string `json:"to_address,omitempty"`
			TokenID        *string `json:"token_id,omitempty"`
			Amount         *string `json:"amount,omitempty"`
			TokenChainID   *uint64 `json:"token_chain_id,omitempty"`
			TokenClass     *string `json:"token_class,omitempty"`
			TokenName      *string `json:"token_name,omitempty"`
			TokenSymbol    *string `json:"token_symbol,omitempty"`
			TokenDecimals  *uint64 `json:"token_decimals,omitempty"`
			Status         string  `json:"status"`
			Kind           string  `json:"kind"`
		} `json:"withdrawals"`
	} `json:"result"`
}

type WithdrawalHistoryResponse struct {
	Address     common.Address                `json:"address"`
	ChainID     uint64                        `json:"source_chain_id"`
	Withdrawals []*dbmodels.WithdrawalDetails `json:"withdrawals"`
}

// GetWithdrawalHistory godoc
//
//	@Summary		Get withdrawals history related to an address
//	@Description	Returns a list of withdrawals of a chain id initiated by an Ethereum address or proven by it.
//	@Description	All withdrawals are fetched which are no older than 90 days.
//	@Description	Only non-finalized withdrawals older than that are fetched.
//	@Description	Withdrawals are fetched in latest-first order.
//
//	@Tags			withdrawals
//	@Accept			json
//	@Produce		json
//	@Param			address		query		string	true	"Ethereum address"
//	@Param			chain_id	query		string	true	"Chain id of withdrawal initiation"
//	@Success		200			{object}	EnvelopedWithdrawalHistoryResponse
//	@Failure		400			{object}	response.ErrorEnvelope
//	@Failure		500			{object}	response.ErrorEnvelope
//	@Router			/withdrawal-history [get]
func (h *WithdrawalsHandler) GetWithdrawalHistory(w http.ResponseWriter, r *http.Request) {
	queryAddress := r.URL.Query().Get("address")
	queryChainID := r.URL.Query().Get("chain_id")

	if queryAddress == "" || !common.IsHexAddress(queryAddress) {
		response.SendErrorResponse(w, http.StatusBadRequest, "address query param missing or malformed")
		return
	}

	if queryChainID == "" {
		response.SendErrorResponse(w, http.StatusBadRequest, "chain id query param missing")
		return
	}

	addr := common.HexToAddress(queryAddress)

	chainIDInt, err := strconv.Atoi(queryChainID)
	if err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, "chain id query param malformed")
		return
	}

	chainID := uint64(chainIDInt)

	currentTime := time.Now()
	ninetyDays := 90 * 24 * time.Hour
	sinceTime := currentTime.Add(-ninetyDays)

	dbResult, err := h.db.WithdrawalHistory(r.Context(), addr, chainID, uint64(sinceTime.Unix()))
	if err != nil {
		response.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = response.SendJSON(
		w,
		http.StatusOK,
		"Withdrawal history fetched successfully",
		WithdrawalHistoryResponse{
			Address:     addr,
			ChainID:     chainID,
			Withdrawals: dbResult,
		},
	)
	if err != nil {
		response.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}
