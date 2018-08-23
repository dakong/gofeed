package appcli

import (
	iocsv "github.com/dakong/gofeed/iocsv"
	rss "github.com/dakong/gofeed/rss"
)

type feed struct {
	index int
	title string
	link  string
}

// State of the application. Stores the users list of feeds, and the current
// feed that the user is browser
type State struct {
	currentFeed rss.Rss
	feeds       []feed
}

// NewState creates an empty state
func NewState() *State {
	return &State{
		currentFeed: rss.Rss{},
		feeds:       []feed{},
	}
}

// SetFeedList sets the state of feeds
func (s *State) SetFeedList() {
	csv := iocsv.NewFile("data.csv")
	records := csv.Read()
	feedList := make([]feed, 0)

	for index, record := range records {
		feedItem := feed{
			index: index,
			title: record[0],
			link:  record[1],
		}
		feedList = append(feedList, feedItem)
	}

	s.feeds = feedList
}

// SetCurrentFeed sets the state of the currently selected feed
func (s *State) SetCurrentFeed(feed rss.Rss) {
	s.currentFeed = feed
}
