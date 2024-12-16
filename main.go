package main

import (
	"fmt"

	"github.com/deofex/selfupdatetest/selfupdatest/selfupdate"
)

var version string

func main() {
	fmt.Printf("Binary version: %s", version)
	err := selfupdate.SelfUpdate(version)
	if err != nil {
		fmt.Printf("Unable to update, using old version: %v\n", err)
	}
	fmt.Println("En hier komt de kubectl magie de we al hebben")
}
