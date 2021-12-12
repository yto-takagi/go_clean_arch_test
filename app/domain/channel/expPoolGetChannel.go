package domain

import (
	"go_clean_arch_test/app/domain"
)

type ExpPoolGetChannel struct {
	expPool domain.ExpPool
	err     error
}

// constructor
func NewExpPoolGetChannel(expPool domain.ExpPool, err error) (*ExpPoolGetChannel, error) {
	expPoolGetChannel := &ExpPoolGetChannel{
		expPool: expPool,
		err:     err,
	}

	return expPoolGetChannel, nil
}

// setter
func (expPoolGetChannel *ExpPoolGetChannel) Set(expPool domain.ExpPool, err error) error {
	expPoolGetChannel.expPool = expPool
	expPoolGetChannel.err = err

	return nil
}

// getter
func (expPoolGetChannel *ExpPoolGetChannel) GetExpPool() domain.ExpPool {
	return expPoolGetChannel.expPool
}

func (expPoolGetChannel *ExpPoolGetChannel) GetErr() error {
	return expPoolGetChannel.err
}
