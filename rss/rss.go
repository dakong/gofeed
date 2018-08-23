package rss

// Article of the RSS element
type Article struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

// Channel tag in RSS feed
type Channel struct {
	Title    string    `xml:"title"`
	Link     string    `xml:"link"`
	Articles []Article `xml:"item"`
}

// Rss ...
type Rss struct {
	Channel Channel `xml:"channel"`
}
