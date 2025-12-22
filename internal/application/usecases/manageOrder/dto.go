package manageOrder

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/bank"
)

type DTO struct {
	Address  string
	BankName string
	Card     bank.Card

	UserUUID     uuid.UUID
	CartItemUUID uuid.UUID
}
