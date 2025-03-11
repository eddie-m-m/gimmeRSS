package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func main() {

	url := "https://feeds.bbci.co.uk/news/world/rss.xml"

	bbcFeed, err := http.Get(url)
	if err != nil {
		fmt.Println("error in fetching url: %v", err)
	}
	defer bbcFeed.Body.Close()

	body, err := io.ReadAll(bbcFeed.Body)
	if err != nil {
		fmt.Println("error while reading feed: %v", err)
	}

	var rss RSS
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		fmt.Println("error while parsing rss: %v", err)
	}

	fmt.Println("Feed: ", rss.Channel.Title)
	for _, item := range rss.Channel.Items {
		fmt.Println("Item: ", item.Title)
	}

}
