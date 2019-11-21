package main

import (
	"fmt"
	"time"
)

type subcommand string

func (subcmd subcommand) Sync() {
	//func main() { //for testing purposes
	fmt.Println("Performing syncing of the Tree ...")

	kits := [4]string{
		"core-kit",
		"python-kit",
		"python-modules-kit",
		"nokit",
	}

	for _, kit := range kits {
		go syncKit(kit)
	}
	time.Sleep(time.Second * 3)
	fmt.Println("Syncing of the Tree [SUCCESS]")
}

func syncKit(kitName string) {
	fmt.Printf("Syncing kit %s ...\n", kitName)
}

// Function exported as symbol named "Function"
var Syncer subcommand
