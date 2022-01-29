package main

import (
	"github.com/Trepka/bookslib/internal/config"
	sv "github.com/Trepka/bookslib/internal/platform/server"
)

func main() {
	cfg := config.PrepareConfig()
	sv.RunServer(cfg.DbConf)
}
