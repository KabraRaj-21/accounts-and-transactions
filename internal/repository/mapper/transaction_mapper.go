package mapper

import (
	"fmt"
	"strconv"
	"transaction/internal/entity"
	"transaction/internal/errors/tserror"
	"transaction/internal/repository/types/record"
)

func MapTransactionRecordToEntity(req *record.Transaction) *entity.Transaction {
	return &entity.Transaction{
		Id:            fmt.Sprintf("%v", req.ID),
		OperationType: entity.OperationType(req.OperationType),
		AccountId:     fmt.Sprintf("%v", req.AccountID),
		Amount:        req.Amount,
		Timestamp:     req.EventTimestamp,
	}
}

func MapTransactionEntityToRecord(req *entity.Transaction) (*record.Transaction, error) {
	// id is not mapped here
	accountId, err := strconv.Atoi(req.AccountId)
	if err != nil {
		return nil, tserror.Wrap(tserror.ErrorType_INVALID_REQUEST, "account id is invalid", err)
	}
	rec := &record.Transaction{
		AccountID:      uint(accountId),
		OperationType:  int(req.OperationType),
		Amount:         req.Amount,
		EventTimestamp: req.Timestamp,
	}
	return rec, nil
}

func GetTransactionRecordIdFromEntity(req *entity.Transaction) (uint, error) {
	res, err := strconv.Atoi(req.Id)
	if err != nil {
		return 0, err
	}
	return uint(res), nil
}
