package main

import (
	"fmt"
	"github.com/rojack96/vierno/routes"
)

func main() {
	fmt.Println("partiti")
	r := routes.SetupRouter()
	err := r.Run(":4788")
	if err != nil {
		return
	}
}
