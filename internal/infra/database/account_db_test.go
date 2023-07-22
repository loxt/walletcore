package database

import (
	"database/sql"
	"github.com/loxt/walletcore/internal/domain/entity"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	AccountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)

	s.db = db
	s.db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at timestamp)")
	s.db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at timestamp)")
	s.AccountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("John Doe", "j@j.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}
