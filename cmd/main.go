package main

import (
	"dev/hak/internal/app"
	"dev/hak/internal/config"
	"log"
)

var cfg config.Config

func init() {
	err := cfg.InitCfg()
	if err != nil {
		panic(err)
	}
}

func main() {
	err := app.Run(cfg)
	if err != nil {
		log.Println(err)
	}
}
