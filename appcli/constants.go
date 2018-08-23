package appcli

import color "github.com/fatih/color"

const (
	cQuit         string = "CTRL + c to quit"
	cInvalidInput string = "There was an issue with the input given"
	cOpenBrowser  string = "Would you like to open in browser? [y/n] "
	cSelectItem   string = "Select an item (h to go home): "
	cSelectFeed   string = "Select a feed: "
	cSaved        string = "Entry was successfully saved"
	cNotSaved     string = "Could not save this entry"
	cOutOfBounds  string = "Selected number is out of bounds"
	cDecodeError  string = "An error occured formatting the feed"
)

var quitCLI = color.RedString(cQuit)
var invalidInputCLI = color.RedString(cInvalidInput)
var openBrowserCLI = color.CyanString(cOpenBrowser)
var selectItemCLI = color.CyanString(cSelectItem)
var selectFeedCLI = color.CyanString(cSelectFeed)
var entrySavedCLI = color.GreenString(cSaved)
var entryNotSavedCLI = color.RedString(cNotSaved)
var outOfBoundsCLI = color.RedString(cOutOfBounds)
var decodeErr = color.RedString(cDecodeError)
