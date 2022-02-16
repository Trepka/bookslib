package main

import (
	"github.com/Trepka/bookslib/internal/config"
	"github.com/Trepka/bookslib/internal/logger"
	server "github.com/Trepka/bookslib/internal/platform/server"
)

func main() {
	cfg := config.PrepareConfig()
	logger := logger.New(cfg.LogConf)
	server.RunServer(cfg.DbConf, logger)
}
