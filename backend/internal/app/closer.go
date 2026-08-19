package app

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrNilFunc = errors.New("close func is nil")
)

type closeFunc func(context.Context) error

type closer struct {
	funcs []closeFunc
	mu sync.Mutex
}

func (c *closer) add(fc closeFunc) error {
	if fc == nil {
		return ErrNilFunc
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.funcs = append(c.funcs, fc)
	return nil
}

func (c *closer) closeAll(ctx context.Context) error {
	c.mu.Lock()
	funcs := make([]closeFunc, len(c.funcs))
	copy(funcs, c.funcs)
	c.funcs = []closeFunc{}
	c.mu.Unlock()

	var errs []error
	for _, fc := range funcs {
		if err := fc(ctx); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return errors.Join(errs...)
	}

	return nil
}

var globalCloser *closer

func init() {
	globalCloser = &closer{
		funcs: []closeFunc{},
	}
}

func Add(fc closeFunc) error {
	return globalCloser.add(fc)
}

func CloseAll(ctx context.Context) error {
	return globalCloser.closeAll(ctx)
}