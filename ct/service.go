package ct

import (
	"context"

	"github.com/cockroachdb/errors"
)

type Service interface {
	Deposit(ctx context.Context, acc string, amount uint) (int, error)
	Withdraw(ctx context.Context, acc string, amount uint) (int, error)
	Balance(ctx context.Context, acc string) (int, error)
}

type Repo interface {
	GetTransactions(ctx context.Context, acc string) ([]Transaction, error)
	InsertTransaction(ctx context.Context, tx Transaction) error
	NewAccount(ctx context.Context, acc string, level Restriction) error
	SetRestriction(ctx context.Context, acc string, level Restriction) error
}

func NewService(r Repo) Service {
	return &service{r}
}

type service struct {
	repo Repo
}

func (s *service) Deposit(ctx context.Context, acc string, amount uint) (int, error) {
	return 0, errors.New("not implemented")
}

func (s *service) Withdraw(ctx context.Context, acc string, amount uint) (int, error) {
	return 0, errors.New("not implemented")
}

func (s *service) Balance(ctx context.Context, acc string) (int, error) {
	return 0, errors.New("not implemented")
}
