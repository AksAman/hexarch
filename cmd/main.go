// This will be the entry point for our application
// orchestrating the startup of the application and
// dependency injection etc.

package main

import (
	"fmt"

	"github.com/AksAman/hexarch/internal/adpaters/core/arithmetic"
	"github.com/AksAman/hexarch/internal/ports"
)

func main() {

	// ports
	var corePort ports.ArithmeticPort

	corePort = arithmetic.NewAdapter()

	res, err := corePort.Division(5, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("res: %v\n", res)
}
