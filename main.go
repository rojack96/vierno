package main

import (
	"github.com/rojack96/vierno/routes"
)

func main() {
	r := routes.SetupRouter()
	err := r.Run(":4788")
	if err != nil {
		return
	}
}
