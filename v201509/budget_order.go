package v201509

import (
//	"encoding/xml"
//	"fmt"
)

type BudgetOrderService struct {
	Auth
}

func NewBudgetOrderService(auth *Auth) *BudgetOrderService {
	return &BudgetOrderService{Auth: *auth}
}
