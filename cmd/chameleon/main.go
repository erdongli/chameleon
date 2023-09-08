package main

import (
	"flag"
	"log"

	"github.com/erdongli/chameleon/internal/dns"
	"github.com/erdongli/chameleon/internal/runner"
)

var (
	username = flag.String("u", "", "DDNS username")
	password = flag.String("p", "", "DDNS password")
	hostname = flag.String("h", "", "hostname to update")
)

func main() {
	flag.Parse()
	if *username == "" || *password == "" || *hostname == "" {
		log.Fatal("missing username/passwod/hostname")
	}

	runner.Run(dns.New(*username, *password, *hostname))
}
