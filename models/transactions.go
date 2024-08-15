package models

import "time"

const (
	TransactionEarning = iota
	TransactionCost
)

type Transaction struct {
	ID       int
	Name     string
	Date     time.Time
	Amount   uint
	Type     uint
	PocketID int
	UserID   int
}
