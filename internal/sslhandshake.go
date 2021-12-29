package internal

import (
	"context"
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

func (s *SSLHandshake) EstablishTcp() (*net.Conn, int64, error) {
	startTime := time.Now()

	conn, err := net.DialTimeout("tcp", s.Config.Endpoint, time.Duration(s.Config.Timeout)*time.Millisecond)
	if err != nil {
		return nil, 0, err
	}

	return &conn, time.Since(startTime).Milliseconds(), nil
}

func (s *SSLHandshake) DoHandshake() (int64, uint16, int64, error) {
	tlsConfig := tls.Config{
		InsecureSkipVerify: true,
	}

	tcpConn, tcpTime, tcpErr := s.EstablishTcp()
	if tcpErr != nil {
		return tcpTime, 0, 0, tcpErr
	}

	startTime := time.Now()
	tlsConn := tls.Client(*tcpConn, &tlsConfig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.Config.Timeout)*time.Millisecond)
	defer cancel()

	tlsErr := tlsConn.HandshakeContext(ctx)
	if tlsErr != nil {
		return tcpTime, 0, 0, tlsErr
	}
	tlsVersion := tlsConn.ConnectionState().Version

	tlsConn.Close()

	return tcpTime, tlsVersion, time.Since(startTime).Milliseconds(), nil
}

func (s *SSLHandshake) Start() {
	var (
		samples   []int64
		errMsg    string
		totalTime int64
	)

	fmt.Printf("Starting %s v%s ( %s ) at %s\n",
		s.Metadata.Name,
		s.Metadata.Version,
		s.Metadata.URL,
		time.Now().Format(time.RFC1123))

	for {
		tcpTime, tlsVersion, handshakeTime, err := s.DoHandshake()
		totalTime = tcpTime + handshakeTime

		s.State.Total++

		tlsVersionName := "n/a"
		switch tlsVersion {
		case tls.VersionTLS10:
			tlsVersionName = "1.0"
		case tls.VersionTLS11:
			tlsVersionName = "1.1"
		case tls.VersionTLS12:
			tlsVersionName = "1.2"
		case tls.VersionTLS13:
			tlsVersionName = "1.3"
		}

		if err != nil {
			s.State.Failed++
			if err, ok := err.(net.Error); ok && err.Timeout() {
				errMsg = "timeout exceeded"
			} else {
				errMsg = err.Error()
			}
			fmt.Printf("SSL handshake with %s: tcp=%dms handshake=%dms version=%s total=%dms error=%s\n", s.Config.Endpoint, tcpTime, handshakeTime, tlsVersionName, totalTime, errMsg)
		} else {
			s.State.Finished++
			samples = append(samples, handshakeTime)
			if handshakeTime > s.State.Max {
				s.State.Max = handshakeTime
			}
			if s.State.Min == 0 || handshakeTime < s.State.Min {
				s.State.Min = handshakeTime
			}
			s.State.Avg = Sum(samples) / s.State.Finished
			fmt.Printf("SSL handshake with %s: tcp=%dms handshake=%dms version=%s total=%dms\n", s.Config.Endpoint, tcpTime, handshakeTime, tlsVersionName, totalTime)
		}

		if s.Config.StopCount != 0 && s.State.Finished == s.Config.StopCount {
			s.ShowStatistics()
		}

		time.Sleep(time.Duration(s.Config.Interval) * time.Millisecond)
		s.State.Elapsed = s.State.Total * s.Config.Interval
	}
}
