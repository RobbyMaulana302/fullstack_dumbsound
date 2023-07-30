package transactiondto

type TransactionRequest struct {
	Status string `json:"status" form:"status" validate:"required"`
}
