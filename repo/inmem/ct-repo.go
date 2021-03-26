package inmem

import (
	"context"
	"sync"

	"github.com/cockroachdb/errors"
	"github.com/nacobas/credit-transfer/ct"
)

type CTRepo struct {
	mtx  sync.RWMutex
	data map[string]*dao
}

func (r *CTRepo) NewAccount(ctx context.Context, acc string, level ct.Restriction) error {

	r.mtx.Lock()
	defer r.mtx.Unlock()

	_, ok := r.data[acc]
	if ok {
		return errors.Newf("Account allready exists %s", acc)
	}

	r.data[acc] = &dao{res: level}

	return nil
}

func (r *CTRepo) SetRestriction(ctx context.Context, acc string, level ct.Restriction) error {

	r.mtx.Lock()
	defer r.mtx.Unlock()

	dao, ok := r.data[acc]
	if !ok {
		return errors.Newf("Account not found %s", acc)
	}

	dao.res = level

	return nil
}

func (r *CTRepo) GetTransactions(ctx context.Context, acc string) ([]ct.Transaction, error) {

	r.mtx.RLock()
	defer r.mtx.RUnlock()

	dao, ok := r.data[acc]
	if !ok {
		return nil, errors.Newf("Account not found %s", acc)
	}

	if dao.res < ct.AllowRead {
		return nil, errors.Newf("Restricted by level: %d", dao.res)
	}

	return toTransactions(acc, dao.txs)
}

func (r *CTRepo) InsertTransaction(ctx context.Context, ctt ct.Transaction) error {

	r.mtx.Lock()
	defer r.mtx.Unlock()

	dao, ok := r.data[ctt.Account()]
	if !ok {
		return errors.Newf("Account not found %s", ctt.Account())
	}

	if len(dao.txs) != int(ctt.Number()) {
		return errors.Newf("Unaccepted trasaction number: %d, latest: %d", ctt.Number(), len(dao.txs)-1)
	}

	if ct.TransactionType(dao.res) < ctt.Type() {
		return errors.Newf("Restricted by level: %d", dao.res)
	}

	dao.txs = append(dao.txs, tx{ctt.Type(), ctt.Amount()})

	return nil
}

type dao struct {
	res ct.Restriction
	txs []tx
}

type tx struct {
	txType ct.TransactionType
	amount uint
}

func toTransactions(acc string, txs []tx) ([]ct.Transaction, error) {

	ret := make([]ct.Transaction, len(txs))

	for i, tx := range txs {
		switch tx.txType {
		case ct.Debit:
			ret[i] = ct.DebitTx(acc, uint(i), tx.amount)
		case ct.Credit:
			ret[i] = ct.DebitTx(acc, uint(i), tx.amount)
		default:
			return nil, errors.Newf("Unknown TransactionType: %d", tx.txType)
		}
	}
	return ret, nil

}
