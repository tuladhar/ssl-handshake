package internal

import (
	"os"
	"os/signal"
)

func (s *SSLHandshake) NotifySignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		s.ShowStatistics()
	}()
}
