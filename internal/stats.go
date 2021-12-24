package internal

import (
	"fmt"
	"os"
)

func (s *SSLHandshake) ShowStatistics() {
	fmt.Println()
	fmt.Printf("--- %s handshake statistics ---\n", s.Config.Endpoint)
	fmt.Printf("%d handshake sent, established: %d, failed: %d\n", s.State.Total, s.State.Success, s.State.Failed)
	fmt.Printf("rtt max: %d ms, avg: %d ms, min: %d ms\n", s.State.Max, s.State.Avg, s.State.Min)
	os.Exit(0)
}
