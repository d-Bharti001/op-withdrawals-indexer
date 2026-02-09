package withdrawals

import (
	"net/http"
	"strconv"
	"time"

	"op-withdrawals-indexer/internal/api/dbmodels"
	"op-withdrawals-indexer/internal/api/handlers/response"

	"github.com/ethereum/go-ethereum/common"
)

type WithdrawalHistoryResponse struct {
	Address     common.Address                `json:"address"`
	ChainID     uint64                        `json:"source_chain_id"`
	Withdrawals []*dbmodels.WithdrawalDetails `json:"withdrawals"`
}

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
		dbResult,
	)
	if err != nil {
		response.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}
