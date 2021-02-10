package main

import (
	"log"

	"github.com/go-delve/delve/pkg/config"
	"https://github.com/shohei1021/nba-batch-go/pkg/config"
)

const (
	Log = "var/log"
)

func main() {
	if err := config.Load("config.toml"); err != nil {
		log.Fatal(err)
	}
}
