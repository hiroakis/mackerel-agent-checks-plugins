package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

const (
	OK       = 0
	WARNING  = 1
	CRITICAL = 2
	UNKNOWN  = 3
)

func main() {
	optHost := flag.String("host", "localhost", "Hostname")
	optPort := flag.String("port", "0", "Port")
	optProtocol := flag.String("protocol", "tcp", "tcp or udp")
	optLevel := flag.String("level", "warn", "warn of crit")
	flag.Parse()

	if *optPort == "0" {
		fmt.Println("-port is required")
		os.Exit(UNKNOWN)
	}

	var msg string
	target := fmt.Sprintf("%s:%s", *optHost, *optPort)
	_, err := net.Dial(*optProtocol, target)
	if err != nil {
		switch *optLevel {
		case "crit":
			msg = fmt.Sprintf("CRITICAL: %s://%s closed", *optProtocol, target)
			fmt.Println(msg)
			os.Exit(CRITICAL)
		case "warn":
			msg = fmt.Sprintf("WARNING: %s://%s closed", *optProtocol, target)
			fmt.Println(msg)
			os.Exit(WARNING)
		}
	}
	msg = fmt.Sprintf("OK: %s://%s open", *optProtocol, target)
	fmt.Println(msg)
	os.Exit(OK)
}
