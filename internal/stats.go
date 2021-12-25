package internal

import (
	"fmt"
	"os"
)

func (s *SSLHandshake) ShowStatistics() {
	fmt.Println()
	fmt.Printf("--- %s handshake statistics ---\n", s.Config.Endpoint)
	fmt.Printf("%d handshakes sent, finish: %d, fail: %d, time: %dms\n", s.State.Total, s.State.Finished, s.State.Failed, s.State.Elapsed)
	fmt.Printf("rtt min/avg/max: %d/%d/%d ms\n", s.State.Min, s.State.Avg, s.State.Max)
	os.Exit(0)
}
