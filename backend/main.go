package main

import (
	"github.com/rojack96/vierno/config"
	"github.com/rojack96/vierno/routes"
)

func main() {

	cfg, err := config.ReadViernoConfig()
	if err != nil {
		// TODO inserire logger
		panic("Error reading configuration: " + err.Error())
	}

	r := routes.SetupRouter(cfg)
	if err = r.Run(":" + cfg.Vierno.Port); err != nil {
		return
	}

}
