package dto

type NewAccountResponse struct {
	AccountId string `json:"account_id"`
}

type NewTransactionResponse struct {
	TransactionId string `json:"transaction_id"`
	Amount        string `json:"amount"`
}
