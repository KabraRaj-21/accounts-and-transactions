package mapper

import (
	"fmt"
	"strconv"
	"transaction/internal/entity"
	"transaction/internal/repository/types/record"
)

func MapAccountRecordToEntity(req *record.Account) *entity.Account {
	return &entity.Account{
		Id:             fmt.Sprintf("%v", req.ID),
		DocumentNumber: req.DocumentNumber,
		Balance:        req.Balance,
	}
}

func MapAccountEntityToRecord(req *entity.Account) *record.Account {
	// id is not mapped here
	return &record.Account{
		DocumentNumber: req.DocumentNumber,
		Balance:        req.Balance,
	}
}

func GetAccountRecordIdFromEntity(req *entity.Account) (uint, error) {
	res, err := strconv.Atoi(req.Id)
	if err != nil {
		return 0, err
	}
	return uint(res), nil
}
