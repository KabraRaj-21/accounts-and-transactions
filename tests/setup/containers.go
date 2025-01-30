package setup

import (
	"context"

	"github.com/Swiggy/grill"
	"github.com/Swiggy/grill/pkg/grillmysql"
)

type TestEnvironment struct {
	MySQL *grillmysql.Mysql
}

func NewTestEnvironment() *TestEnvironment {
	return &TestEnvironment{
		MySQL: &grillmysql.Mysql{},
	}
}

func (t *TestEnvironment) StartAll() error {
	return grill.StartAll(context.Background(), t.MySQL)
}

func (t *TestEnvironment) StopAll() error {
	return grill.StopAll(context.Background(), t.MySQL)
}

func (t *TestEnvironment) SetupTables() grill.StubFunc {
	return func() error {
		grillStubs := []grill.Stub{
			t.MySQL.CreateTable(CreateAccountsTable),
			t.MySQL.CreateTable(CreateTransactionsTable),
		}

		for _, stub := range grillStubs {
			if err := stub.Stub(); err != nil {
				return err
			}
		}
		return nil
	}
}
func (t *TestEnvironment) CleanTables() grill.CleanerFunc {
	return func() error {
		cleaners := []grill.Cleaner{
			t.MySQL.DeleteTable("transactions"),
			t.MySQL.DeleteTable("accounts"),
		}

		for _, cleaner := range cleaners {
			if err := cleaner.Clean(); err != nil {
				return err
			}
		}
		return nil
	}
}
