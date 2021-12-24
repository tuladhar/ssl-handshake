package internal

import (
	"flag"
	"fmt"
	"os"
)

func (s *SSLHandshake) SetUsage() {
	flag.Usage = func() {
		fmt.Printf("%s v%s ( %s )\n\n", s.Metadata.Name, s.Metadata.Version, s.Metadata.URL)
		fmt.Println(s.Metadata.Description)
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Printf("  %s [options] hostname:port\n", s.Metadata.Name)
		fmt.Println()
		fmt.Println("Options:")
		fmt.Printf("  -t <timeout>   SSL handshake timeout in milliseconds (default: %d)\n", s.Config.Timeout)
		fmt.Printf("  -c <count>     Stop after <count> SSL handshakes (default: %d)\n", s.Config.StopCount)
		fmt.Printf("  -i <interval>  Wait <interval> milliseconds between SSL handshakes (default: %d)\n", s.Config.Interval)
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Printf("  %s tuladhar.github.com:443\n", s.Metadata.Name)
		fmt.Printf("  %s -c 3 -i 500 tuladhar.github.com:443\n", s.Metadata.Name)
		fmt.Printf("  %s -t 500 imap.gmail.com:993\n", s.Metadata.Name)
	}
}

func (s *SSLHandshake) ParseFlags() {
	flag.Int64Var(&s.Config.Timeout, "t", s.Config.Timeout, "")
	flag.Int64Var(&s.Config.StopCount, "c", s.Config.StopCount, "")
	flag.Int64Var(&s.Config.Interval, "i", s.Config.Interval, "")

	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	s.Config.Endpoint = flag.Args()[0]
}

func (s *SSLHandshake) InitCLI() {
	s.SetUsage()
	s.ParseFlags()
}
