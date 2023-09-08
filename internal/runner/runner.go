package runner

import (
	"log"
	"time"

	"github.com/erdongli/chameleon/internal/dns"
	"github.com/erdongli/chameleon/internal/ip"
)

const (
	interval = 5 * time.Minute
)

func Run(updater *dns.Updater) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		ip, err := ip.Get()
		if err != nil {
			log.Println(err)
			continue
		}

		ok, err := updater.Update(ip)
		if err != nil {
			log.Println(err)
			continue
		}
		if !ok {
			log.Printf("ip not changed: %s", ip)
			continue
		}

		log.Printf("DDNS updated to %s", ip)
	}
}
