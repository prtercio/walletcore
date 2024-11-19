package database

import (
	"database/sql"
	"testing"

	"github.com/prtercio/fs-ws-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory")
	s.Nil(err)
	s.db = db

	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	db.Exec("Create table transations (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")

	client, err := entity.NewClient("Ben Bao", "ben@bao.com")
	s.Nil(err)
	s.client = client

	client2, err := entity.NewClient("Dan Tau", "dan@tau.com")
	s.Nil(err)
	s.client2 = client2

	accountFrom := entity.NewAccount(s.client)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom
	accountTo := entity.NewAccount(s.client2)
	accountTo.Balance = 1000
	s.accountTo = accountTo

	s.transactionDB = *NewTransaction(db)
}

func TesTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
