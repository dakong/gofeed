package main

import (
	"fmt"
	"os"

	"github.com/dakong/gofeed/appcli"
	flag "github.com/ogier/pflag"
)

var app *appcli.App
var cmd *command
var commandMap map[string]*flag.FlagSet

var nameFlag string
var urlAddFlag string
var urlSearchFlag string

func handleSearchCommand(searchCommand *flag.FlagSet) {
	// Handle search commands
	url := searchCommand.Lookup("url")

	if url != nil {
		app.LookupFeed(url.Value.String())
	}
}

func handleAddCommand(addCommand *flag.FlagSet) {
	url := addCommand.Lookup("url")
	name := addCommand.Lookup("name")

	if url != nil && name != nil {
		appcli.StoreFeed(url.Value.String(), name.Value.String())
	}
}

func init() {
	app = appcli.InitializeApp()
	cmd = initCommands()
}

func main() {

	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("Command is required, usage of gofeed:")
		printHelp()

		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		cmd.add.flagSet.Parse(os.Args[2:])

	case "search":
		cmd.search.flagSet.Parse(os.Args[2:])

	case "home":
		app.ReadFeeds()

	default:
		fmt.Println("Invalid command, usage of gofeed:")
		printHelp()

		os.Exit(1)
	}

	if cmd.add.flagSet.Parsed() {
		handleAddCommand(cmd.add.flagSet)

	} else if cmd.search.flagSet.Parsed() {
		handleSearchCommand(cmd.search.flagSet)

	}
}
