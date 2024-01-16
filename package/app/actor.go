package app

import (
	"context"
	"errors"
	"sync"

	"github.com/samber/lo"
)

type Actors map[string]Actor

type Actor interface {
	Start(context.Context) (bool, error)
	Stop(context.Context) (bool, error)
	Spawn(string, Actor) Actor
	SpawnAll(Actors) Actor
}

type BaseActor struct {
	ctr      uint
	Children Actors
	Lock     *sync.Mutex
	closer   chan struct{}
}

func NewActor() *BaseActor {
	return &BaseActor{
		Children: make(Actors),
		Lock:     &sync.Mutex{},
	}
}
func (a *BaseActor) Spawn(k string, v Actor) Actor {
	if _, ok := a.Children[k]; !ok {
		a.Children[k] = v
	}
	return a
}
func (a *BaseActor) SpawnAll(m Actors) Actor {
	for k, v := range m {
		a.Spawn(k, v)
	}
	return a
}
func (a *BaseActor) Start(ctx context.Context) (bool, error) {
	a.Lock.Lock()
	defer a.Lock.Unlock()
	return a.BaseStart(ctx)
}
func (a *BaseActor) BaseStart(ctx context.Context) (bool, error) {
	a.ctr++
	if a.ctr > 1 {
		return false, nil
	}
	a.closer = make(chan struct{})
	if len(a.Children) == 0 {
		return false, nil
	}
	wg := sync.WaitGroup{}
	wg.Add(len(a.Children))
	errs := make([]error, len(a.Children))
	for i, c := range lo.Keys(a.Children) {
		go func(i int, k string) {
			defer wg.Done()
			_, errs[i] = a.Children[k].Start(ctx)
		}(i, c)
	}
	wg.Wait()
	return true, errors.Join(errs...)
}
func (a *BaseActor) Stop(ctx context.Context) (bool, error) {
	a.Lock.Lock()
	defer a.Lock.Unlock()
	return a.BaseStop(ctx)
}
func (a *BaseActor) BaseStop(ctx context.Context) (bool, error) {
	a.ctr--
	if a.ctr > 0 {
		return false, nil
	}
	wg := sync.WaitGroup{}
	wg.Add(len(a.Children))
	errs := make([]error, len(a.Children))
	for i, c := range lo.Keys(a.Children) {
		go func(i int, k string) {
			defer wg.Done()
			_, errs[i] = a.Children[k].Stop(ctx)
		}(i, c)
	}
	wg.Wait()
	return true, nil
}
func (a *BaseActor) IsStarted() bool { return a.ctr > 0 }
