package app

import (
	"context"
	"svix-poc/lib/utils"
	"sync"
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
}

func NewActor(children Actors) *BaseActor {
	a := &BaseActor{
		Children: make(Actors),
		Lock:     &sync.Mutex{},
	}
	a.SpawnAll(children)
	return a
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
	if a.ctr++; a.ctr > 1 { // if not first call
		return false, nil
	}
	return true, utils.ForAll(a.Children, func(k string) error {
		_, err := a.Children[k].Start(ctx)
		return err
	})
}
func (a *BaseActor) Stop(ctx context.Context) (bool, error) {
	a.Lock.Lock()
	defer a.Lock.Unlock()
	return a.BaseStop(ctx)
}
func (a *BaseActor) BaseStop(ctx context.Context) (bool, error) {
	if a.ctr--; a.ctr > 0 { // if not last call
		return false, nil
	}
	return true, utils.ForAll(a.Children, func(k string) error {
		_, err := a.Children[k].Stop(ctx)
		return err
	})
}
func (a *BaseActor) IsStarted() bool { return a.ctr > 0 }
