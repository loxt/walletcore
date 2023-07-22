package create_transaction

import (
	"github.com/loxt/walletcore/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "j@j.com")
	account := entity.NewAccount(client)
	account.Deposit(1000)

	client2, _ := entity.NewClient("John Doe2", "j@j2.com")
	account2 := entity.NewAccount(client2)
	account2.Deposit(1000)

	accountMock := &AccountGatewayMock{}
	accountMock.On("FindByID", account.ID).Return(account, nil)
	accountMock.On("FindByID", account2.ID).Return(account2, nil)

	transactionMock := &TransactionGatewayMock{}
	transactionMock.On("Create", mock.Anything).Return(nil)

	input := &CreateTransactionInputDTO{
		AccountIDFrom: account.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	uc := NewCreateTransactionUseCase(transactionMock, accountMock)

	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	accountMock.AssertExpectations(t)
	transactionMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "FindByID", 2)
	transactionMock.AssertNumberOfCalls(t, "Create", 1)
}
