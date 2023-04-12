package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("C1Script\n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	//	args := os.Args
	switch {
	//-------------------------------------------------------------
	case len(os.Args) == 2:
		fmt.Printf("cmd %s\n", os.Args[1])
		fmt.Println("Not")

		//}
		//-------------------------------------------------------------
	default:
		fmt.Printf("Command %s\n", "test")

	}
}
