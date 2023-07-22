package usecase

import (
	"github.com/loxt/walletcore/internal/domain/entity"
	"github.com/loxt/walletcore/internal/domain/repository"
	"time"
)

type CreateClientInputDTO struct {
	Name  string
	Email string
}

type CreateClientOutputDTO struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateClientUseCase struct {
	ClientRepository repository.ClientRepository
}

func NewCreateClientUseCase(clientRepository repository.ClientRepository) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientRepository: clientRepository,
	}
}

func (uc *CreateClientUseCase) Execute(input *CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	err = uc.ClientRepository.Save(client)
	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}, nil
}
