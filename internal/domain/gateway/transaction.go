package gateway

import "github.com/loxt/walletcore/internal/domain/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
