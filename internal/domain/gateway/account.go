package gateway

import "github.com/loxt/walletcore/internal/domain/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
