package appcli

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	utils "github.com/dakong/gofeed/utils"
)

// promptSelectArticle prompts user to select an article
func (app *App) promptSelectArticle() (index int) {
	fmt.Print(selectItemCLI)
	input, err := app.reader.ReadString('\n')

	if utils.HandleError(err, invalidInputCLI, false) {
		return app.promptSelectArticle()
	}
	trimmedInput := utils.Trim(input)

	if trimmedInput == "h" || trimmedInput == "home" {
		app.ReadFeeds()
	}

	index, err = strconv.Atoi(trimmedInput)

	if utils.HandleError(err, invalidInputCLI, false) &&
		app.withinArticleRange(index) {
		app.promptSelectArticle()
	}

	return
}

// promptSelectFeed prompts user to select a news feed
func (app *App) promptSelectFeed() (index int) {
	fmt.Print(selectFeedCLI)
	input, err := app.reader.ReadString('\n')

	if utils.HandleError(err, invalidInputCLI, false) {
		return app.promptSelectFeed()
	}

	trimmedInput := utils.Trim(input)
	index, err = strconv.Atoi(trimmedInput)

	if utils.HandleError(err, invalidInputCLI, false) &&
		app.withinFeedRange(index) {
		app.promptSelectFeed()
	}
	return
}

// promptFeedUrl prompts the user for a url in order to view a RSS feed
func (app *App) promptFeedURL() {
	fmt.Println("Enter a valid RSS feed url")
	input, err := app.reader.ReadString('\n')

	if utils.HandleError(err, invalidInputCLI, false) {
		app.promptFeedURL()
	} else {
		app.LookupFeed(input)
	}
}

// promptOpenInBrowser prompts the user to open the article in a web browser
func (app *App) promptOpenInBrowser(index int, link string) {
	fmt.Print(openBrowserCLI)

	input, err := app.reader.ReadString('\n')
	if utils.HandleError(err, invalidInputCLI, false) {
		app.displayArticleDescription(index)
	} else {
		trimmedInput := utils.Trim(input)
		lowerInput := strings.ToLower(trimmedInput)

		if lowerInput == "y" || lowerInput == "yes" {
			cmd := exec.Command("open", link)
			cmd.Run()
		}

		app.displayFeedArticles()
	}
}
