package database

import (
	"database/sql"
	"github.com/loxt/walletcore/internal/domain/entity"
)

type TransactionDB struct {
	DB *sql.DB
}

func NewTransactionDB(db *sql.DB) *TransactionDB {
	return &TransactionDB{DB: db}
}

func (t *TransactionDB) Create(transaction *entity.Transaction) error {
	_, err := t.DB.Exec("INSERT INTO transactions (id, account_id_from, account_id_to, amount, created_at) VALUES ($1, $2, $3, $4, $5)",
		transaction.ID,
		transaction.AccountFrom.ID,
		transaction.AccountTo.ID,
		transaction.Amount,
		transaction.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
