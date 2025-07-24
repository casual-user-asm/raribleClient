package main

import (
	"github.com/casual-user-asm/raribleClient/config"

	"github.com/casual-user-asm/raribleClient/internal/service"
)

func main() {
	config.LoadEnvFile()
	service.StartServer()
}
