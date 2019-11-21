package main

import (
	"fmt"
	"os"
	"plugin"
)

// Syncer is an interace
type Syncer interface {
	Sync()
}

func main() {
	// determine plugin to load
	subcmd := "help"
	if len(os.Args) == 2 {
		subcmd = os.Args[1]
	}
	var mod string
	switch subcmd {
	case "sync":
		mod = "/usr/share/ego/plugins/sync.so"
	default:
		fmt.Println("don't know that subcommand")
		os.Exit(1)
	}

	plug, err := loadMod(mod)

	// 2. look up a symbol (an exported function or variable)
	// in this case, variable Syncer
	symSyncer, err := plug.Lookup("Syncer")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type Syncer (defined above)
	var syncer Syncer
	syncer, ok := symSyncer.(Syncer)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	// 4. use the module
	syncer.Sync()

}

// load module
func loadMod(mod string) (*plugin.Plugin, error) {
	// 1. open the so file to load the symbols
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return plug, nil
}
