package database

import (
	"database/sql"
	"testing"

	"github.com.br/Soter-Tec/ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
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
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	s.AccountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("john", "j@j.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.AccountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByID() {
	s.db.Exec("insert into clients (id, name, email, created_at) values (?, ?, ?,?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt,
	)
	account := entity.NewAccount(s.client)
	err := s.AccountDB.Save(account)
	s.Nil(err)
	accountDb, err := s.AccountDB.FindByID(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDb.ID)
	s.Equal(account.Client.ID, accountDb.Client.ID)
	s.Equal(account.Balance, accountDb.Balance)
	s.Equal(account.Client.ID, accountDb.Client.ID)
}
