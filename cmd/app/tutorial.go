package app

import (
	"perpustakaan/config"
	"perpustakaan/internal/server"
)

func Start() {
	config.Load()
	server.Run()
}
