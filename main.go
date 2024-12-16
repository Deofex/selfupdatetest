package main

import (
	"fmt"
	"os"
)

func main() {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}

	fmt.Println("Old version at: ", path)
}
