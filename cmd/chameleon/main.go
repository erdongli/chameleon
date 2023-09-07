package main

import (
	"time"

	"github.com/erdongli/chameleon/internal/runner"
)

func main() {
	runner.Run(1 * time.Minute)
}
