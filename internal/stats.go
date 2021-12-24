package internal

import (
	"fmt"
	"os"
)

func (s *SSLHandshake) ShowStatistics() {
	fmt.Println()
	fmt.Printf("--- %s handshake statistics ---\n", s.Config.Endpoint)
	fmt.Printf("%d handshake sent, finished: %d, failed: %d, time: %d ms\n", s.State.Total, s.State.Finished, s.State.Failed, s.State.Elapsed)
	fmt.Printf("rtt max: %d ms, avg: %d ms, min: %d ms\n", s.State.Max, s.State.Avg, s.State.Min)
	os.Exit(0)
}
