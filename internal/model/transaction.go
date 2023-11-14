package model

// TransactionInitiationRequest defines the structure for a transaction initiation request.
type TransactionInitiationRequest struct {
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	RecipientID string  `json:"recipient_id"`
}

// TransactionInitiationResponse defines the structure for a response after initiating a transaction.
type TransactionInitiationResponse struct {
	TransactionID string `json:"transaction_id"`
}
