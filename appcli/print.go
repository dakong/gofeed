package appcli

import (
	"fmt"
	"html"

	"github.com/dakong/gofeed/rss"
	color "github.com/fatih/color"
)

func printFeeds(feeds []feed) {
	// Store Feeds in a map
	for _, feed := range feeds {
		magenta := color.New(color.FgMagenta).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()

		number := green("[", feed.index, "]")
		title := magenta(feed.title)

		fmt.Println(number, title)
	}
}

func printItems(items []rss.Article) {
	fmt.Println()
	for index, item := range items {
		magenta := color.New(color.FgMagenta).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()

		number := green("[", index, "]")
		title := magenta(html.UnescapeString(item.Title))

		fmt.Println(number, title)
	}
	fmt.Println()
}

func printItemDescription(description string) {
	fmt.Printf("\n%s\n\n", color.HiMagentaString(description))
}

func printSaved() {
	fmt.Println(entrySavedCLI)
}

func printSavedFail() {
	fmt.Println(entryNotSavedCLI)
}

func printOutOfBounds() {
	fmt.Println(outOfBoundsCLI)
}
