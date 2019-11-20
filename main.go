package main

import (
	"flag"
	"fmt"
	"os"

	"gitlab.com/liguros/ego/libs/query"
)

type commandArgs struct {
	Sync struct {
		Help bool
		Kits bool
	}
	Query struct {
		Versions string
		Origin   string
	}
}

func main() {

	// create an instance of the struct called scArgs that will be a pointer to it
	// this syntax creates a pointer.. very handy
	scArgs := &commandArgs{}

	// create subcommand sync
	syncCmd := flag.NewFlagSet("sync", flag.ExitOnError)
	syncCmd.BoolVar(&scArgs.Sync.Help, "help", false, "display usage information")
	//syncCmd.BoolVar(&scArgs.Sync.Kits, "kits-only", false, "Do not sync meta-repo, only kits.")

	// create subcommand query
	queryCmd := flag.NewFlagSet("query", flag.ExitOnError)
	queryCmd.StringVar(&scArgs.Query.Origin, "origin", "", "origin")
	queryCmd.StringVar(&scArgs.Query.Versions, "versions", "", "versions")

	// check
	if len(os.Args) < 2 {
		fmt.Println("Basic usage for ego:")
		fmt.Println("supported subcommands are")
		fmt.Println(" - sync")
		fmt.Println(" - query")
		fmt.Println("...")
		os.Exit(1)
	}

	// look for the first argument from command line and dispatch accordingly
	switch os.Args[1] {
	case "sync":
		syncCmd.Parse(os.Args[2:])
		if scArgs.Sync.Help == true {
			syncCmd.PrintDefaults()
			//fmt.Println("Usage for ego sync")
		} else {
			sync()
		}

	case "query":
		queryCmd.Parse(os.Args[2:])
		if scArgs.Query.Versions != "" {
			query.Versions(scArgs.Query.Versions)
		} else if scArgs.Query.Origin != "" {
			query.Origin(scArgs.Query.Origin)
		} else {
			os.Exit(1)
		}

	case "profile":
		profile()

	case "news":
		news()

	case "config":
		config()

	case "boot":
		boot()
	}

}

// Sync document comment
func sync() {
	fmt.Println("Syncing")
}

// Profile document comment
func profile() {
	fmt.Println("profiling...")
}

// News document comment
func news() {
	fmt.Println("news...")
}

// Config document comment
func config() {
	fmt.Println("config...")
}

// Boot document comment
func boot() {
	fmt.Println("boot...")
}
