package appcli

import (
	"bufio"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"os"

	iocsv "github.com/dakong/gofeed/iocsv"
	rss "github.com/dakong/gofeed/rss"
	utils "github.com/dakong/gofeed/utils"
)

// App ...
type App struct {
	reader *bufio.Reader
	state  *State
}

// InitializeApp ...
func InitializeApp() *App {
	return &App{
		reader: bufio.NewReader(os.Stdin),
		state:  NewState(),
	}
}

// withinFeedRange check if the index is within range of the list of feeds
func (app *App) withinFeedRange(index int) bool {
	if len(app.state.feeds) >= index+1 && index > 0 {
		return true
	}
	printOutOfBounds()
	return false
}

// withinFeedRange checks if the index is within range of the list of items
func (app *App) withinArticleRange(index int) bool {
	if len(app.state.currentFeed.Channel.Articles) >= index+1 && index > 0 {
		return true
	}
	printOutOfBounds()
	return false
}

// decodeXMLdata will parse XML data and format it to the rss struct
func (app *App) decodeXMLData(xmlData io.ReadCloser) {
	currentFeed := rss.Rss{}
	decoder := xml.NewDecoder(xmlData)
	err := decoder.Decode(&currentFeed)

	if utils.HandleError(err, decodeErr, false) {
		os.Exit(1)
	}

	app.state.SetCurrentFeed(currentFeed)
}

// selectFeed will look up the feed that the user selected
func (app *App) selectFeed() {
	selectedIndex := app.promptSelectFeed()

	for _, feed := range app.state.feeds {
		if feed.index == selectedIndex {
			app.LookupFeed(feed.link)
			return
		}
	}
}

// selectArticle will show the description of the article that the user selected
func (app *App) selectArticle() {
	// Show the select feed item prompt
	selectedIndex := app.promptSelectArticle()

	// Display the description of the feed item
	app.displayArticleDescription(selectedIndex)
}

// displayItemDescription displays the item's description. Then prompt to open
// the item in the browser
func (app *App) displayArticleDescription(index int) {
	selectedItem := app.state.currentFeed.Channel.Articles[index]
	description := html.UnescapeString(selectedItem.Description)

	printItemDescription(description)
	app.promptOpenInBrowser(index, selectedItem.Link)
}

// displayFeedArticles shows users a list of articles to select
func (app *App) displayFeedArticles() {
	printItems(app.state.currentFeed.Channel.Articles)
	app.selectArticle()
}

// displayFeedList shows user a list of feeds to select
func (app *App) displayFeedList() {
	printFeeds(app.state.feeds)
	app.selectFeed()
}

// LookupFeed will search and return a feed given the url
func (app *App) LookupFeed(feedURL string) {
	// Fetch feed from the url
	resp, err := http.Get(feedURL)

	// When there is an error prompt the user for a valid url
	if utils.HandleError(err, invalidInputCLI, false) {
		app.promptFeedURL()
	} else {
		defer resp.Body.Close()

		app.decodeXMLData(resp.Body)
		app.displayFeedArticles()
	}
}

// StoreFeed will store the feed url and name to a file to the data file
func StoreFeed(feedURL string, feedName string) {
	csv := iocsv.NewFile("data.csv")
	entry := []string{feedName, feedURL}

	if csv.Store(entry) {
		printSaved()
	} else {
		printSavedFail()
	}
}

// ReadFeeds will return the list of stored feeds
func (app *App) ReadFeeds() {
	app.state.SetFeedList()
	app.displayFeedList()
}
