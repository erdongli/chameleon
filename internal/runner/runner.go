package runner

import (
	"log"
	"time"

	"github.com/erdongli/chameleon/internal/ip"
)

func Run(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		ip, err := ip.Get()
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println(ip)
	}
}
