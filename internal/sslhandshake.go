package internal

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

type SSLHandshakeMetadata struct {
	Name        string
	Description string
	Version     string
	URL         string
}

type SSLHandshakeConfig struct {
	Endpoint  string
	Interval  int64
	Timeout   int64
	StopCount int64
}

type SSLHandshakeState struct {
	Total    int64
	Finished int64
	Failed   int64
	Max      int64
	Avg      int64
	Min      int64
	Elapsed  int64
}

type SSLHandshake struct {
	Metadata *SSLHandshakeMetadata
	Config   *SSLHandshakeConfig
	State    *SSLHandshakeState
}

func (s *SSLHandshake) DoHandshake() (time.Duration, error) {
	start := time.Now()
	var end time.Duration

	dialer := net.Dialer{
		Timeout: time.Duration(s.Config.Timeout) * time.Millisecond,
	}
	tlsConfig := tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.DialWithDialer(&dialer, "tcp", s.Config.Endpoint, &tlsConfig)
	if err != nil {
		return end, err
	}
	end = time.Since(start)
	conn.Close()
	return end, nil
}

func (s *SSLHandshake) Start() {
	var (
		ssl_handshake_samples []int64
		ssl_handshake_error   string
		handshake_in_ms       int64
	)

	fmt.Printf("Starting %s v%s ( %s ) at %s\n",
		s.Metadata.Name,
		s.Metadata.Version,
		s.Metadata.URL,
		time.Now().Format(time.RFC1123))

	for {
		handshake_duration, err := s.DoHandshake()
		s.State.Total++

		if err != nil {
			s.State.Failed++
			if err, ok := err.(net.Error); ok && err.Timeout() {
				ssl_handshake_error = fmt.Sprintf("timeout (%d ms) exceeded", s.Config.Timeout)
			} else {
				ssl_handshake_error = err.Error()
			}
			fmt.Printf("SSL handshake with %s: time=%d ms error=%s\n", s.Config.Endpoint, 0, ssl_handshake_error)
		} else {
			s.State.Finished++
			handshake_in_ms = handshake_duration.Milliseconds()
			ssl_handshake_samples = append(ssl_handshake_samples, handshake_in_ms)

			if handshake_in_ms > s.State.Max {
				s.State.Max = handshake_in_ms
			}

			if s.State.Min == 0 || handshake_in_ms < s.State.Min {
				s.State.Min = handshake_in_ms
			}

			s.State.Elapsed = Sum(ssl_handshake_samples)
			s.State.Avg = Sum(ssl_handshake_samples) / s.State.Finished

			fmt.Printf("SSL handshake with %s: time=%d ms\n", s.Config.Endpoint, handshake_in_ms)
		}

		if s.State.Total == s.Config.StopCount {
			s.ShowStatistics()
		}

		time.Sleep(time.Duration(s.Config.Interval) * time.Millisecond)
	}
}
