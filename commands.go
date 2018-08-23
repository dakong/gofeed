package main

import (
	"fmt"

	flag "github.com/ogier/pflag"
)

type addCmd struct {
	flagSet *flag.FlagSet
	name    string

	nameFlag string
	urlFlag  string
}

type searchCmd struct {
	flagSet *flag.FlagSet
	name    string

	urlFlag string
}

type command struct {
	add    *addCmd
	search *searchCmd
}

type commandFlag interface {
	printDefault()
}

func (a addCmd) printDefault() {
	a.flagSet.PrintDefaults()
}

func (s searchCmd) printDefault() {
	s.flagSet.PrintDefaults()
}

func printCmdHelp(cmd commandFlag) {
	cmd.printDefault()
}

func printHelp() {
	fmt.Println("Add:")
	printCmdHelp(cmd.add)

	fmt.Println("Search:")
	printCmdHelp(cmd.search)
}

func initAddCmd() (add *addCmd) {
	const cAdd string = "add"

	add = &addCmd{}
	add.flagSet = flag.NewFlagSet(cAdd, flag.ExitOnError)
	add.name = cAdd

	add.flagSet.StringVarP(&add.nameFlag, "name", "n", "", "Name of the RSS feed")
	add.flagSet.StringVarP(&add.urlFlag, "url", "u", "", "Url of the RSS feed")

	return
}

func initSearchCmd() (search *searchCmd) {
	const cSearch string = "search"

	search = &searchCmd{}
	search.flagSet = flag.NewFlagSet(cSearch, flag.ExitOnError)
	search.name = cSearch

	search.flagSet.StringVarP(&search.urlFlag, "url", "u", "", "Url of the RSS feed")

	return
}

func initCommands() *command {
	addCmd := initAddCmd()
	searchCmd := initSearchCmd()

	return &command{
		add:    addCmd,
		search: searchCmd,
	}
}

// func initializeCommands() map[string]*flag.FlagSet {

// 	addCommand := flag.NewFlagSet(cAdd, flag.ExitOnError)
// 	searchCommand := flag.NewFlagSet(cSearch, flag.ExitOnError)

// 	return map[string]*flag.FlagSet{
// 		cAdd:    addCommand,
// 		cSearch: searchCommand,
// 	}
// }
