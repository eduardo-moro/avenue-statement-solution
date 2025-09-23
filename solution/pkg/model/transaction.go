package model

import "time"

// TransactionMetadata represents the metadata of a transaction.
//
// Fields:
//  Description: The description of the transaction.
//  Source: The source of the transaction.
//  Reference: The reference of the transaction.
type TransactionMetadata struct {
	Description string `json:"description"`
	Source      string `json:"source,omitempty"`
	Reference   string `json:"reference,omitempty"`
}

// TransactionEvent represents a transaction event.
//
// Fields:
//  ID: The unique identifier of the transaction.
//  UserID: The identifier of the user.
//  Account: The account of the transaction.
//  Currency: The currency of the transaction.
//  Type: The type of the transaction.
//  Direction: The direction of the transaction.
//  Amount: The amount of the transaction.
//  Balance: The balance of the account after the transaction.
//  Metadata: The metadata of the transaction.
//  ProcessedAt: The timestamp when the transaction was processed.
//  CreatedAt: The timestamp when the transaction was created.
type TransactionEvent struct {
	ID          string              `json:"id"`
	UserID      string              `json:"user_id"`
	Account     string              `json:"account"`
	Currency    string              `json:"currency"`
	Type        string              `json:"type"`
	Direction   string              `json:"direction"`
	Amount      float64             `json:t"amount"`
	Balance     float64             `json:"balance"`
	Metadata    TransactionMetadata `json:"metadata"`
	ProcessedAt time.Time           `json:"processed_at"`
	CreatedAt   time.Time           `json:"created_at"`
}