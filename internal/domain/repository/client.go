package repository

import "github.com/loxt/walletcore/internal/domain/entity"

type ClientRepository interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
