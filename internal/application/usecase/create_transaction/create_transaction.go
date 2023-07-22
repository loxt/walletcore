package create_transaction

import (
	"github.com/loxt/walletcore/internal/domain/entity"
	"github.com/loxt/walletcore/internal/domain/gateway"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	transactionGateway gateway.TransactionGateway
	accountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		transactionGateway: transactionGateway,
		accountGateway:     accountGateway,
	}
}

func (uc *CreateTransactionUseCase) Execute(input *CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := uc.accountGateway.FindByID(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := uc.accountGateway.FindByID(input.AccountIDTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)

	if err != nil {
		return nil, err
	}

	err = uc.transactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutputDTO{
		ID: transaction.ID,
	}, nil
}
