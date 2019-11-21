package main

import "fmt"

type subcommand string

func (subcmd subcommand) Sync() {
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

}

func syncKit(kitName string) {
	fmt.Printf("Syncing kit %s ...\n", kitName)
}

// Function exported as symbol named "Function"
var Syncer subcommand
