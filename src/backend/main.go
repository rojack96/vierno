package main

import (
	"os"
	"strconv"

	"github.com/rojack96/vierno/config"
	"github.com/rojack96/vierno/routes"
)

func main() {
	devMode, _ := strconv.ParseBool(os.Getenv("DEVMODE"))
	cfg, err := config.ReadViernoConfig(devMode)
	if err != nil {
		// TODO inserire logger
		panic("Error reading configuration: " + err.Error())
	}

	r := routes.SetupRouter(cfg, devMode)
	if err = r.Run(":" + cfg.Vierno.Port); err != nil {
		return
	}

}
