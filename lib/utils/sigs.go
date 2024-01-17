package utils

import (
	"os"
	"os/signal"
)

func WaitSigs(sigs ...os.Signal) <-chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, sigs...)
	return ch
}
