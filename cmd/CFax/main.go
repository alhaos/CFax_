package main

import (
	_ "embed"
	"fmt"
	"github.com/alhaos/cfax/internal/webface"
)

// init start web face service
func init() {

	go webface.Start()
}

func main() {
	fmt.Println("Hello world")
	for {

	}
}
