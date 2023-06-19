package transaction

import ut "github.com/berrylradianh/ecowave-go/modules/usecase/user/transaction"

type TransactionHandler struct {
	transactionUsecase ut.TransactionUsecase
}

func New(transactionUsecase ut.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{
		transactionUsecase,
	}
}
