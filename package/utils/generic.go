package utils

import (
	"errors"
	"sync"

	"github.com/samber/lo"
)

func Apply[T any](t T, opts []func(T)) {
	for _, opt := range opts {
		opt(t)
	}
}

func ForAll[T any](in map[string]T, fn func(string) error) error {
	if len(in) == 0 {
		return nil
	}
	wg := sync.WaitGroup{}
	wg.Add(len(in))
	errs := make([]error, len(in))
	for i, c := range lo.Keys(in) {
		go func(i int, k string) {
			defer wg.Done()
			errs[i] = fn(k)
		}(i, c)
	}
	wg.Wait()
	return errors.Join(errs...)
}
