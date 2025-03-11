package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Feed struct {
	FeedName string `json:"name"`
	FeedURL  string `json:"url"`
}

func AddFeeds() {

	var feeds []Feed
	var jsonFileName = "feeds.json"

	jsonFile, err := os.ReadFile(jsonFileName)
	if os.IsNotExist(err) {
		os.Create(jsonFileName)
	} else {
		err = json.Unmarshal(jsonFile, &feeds)
		if err != nil {
			err = fmt.Errorf("error while unmarshaling JSON: %v", err)
			fmt.Println(err)
		}
	}

	for {
		var feed Feed
		fmt.Print(">>>Enter a shorthand name for the feed:")
		fmt.Scanln(&feed.FeedName)

		fmt.Print(">>>Enter the feed's URL:")
		fmt.Scanln(&feed.FeedURL)

		feeds = append(feeds, feed)

		var input string
		fmt.Print("Add another? y[es] or n[o]")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(input)
		switch input {
		case "y":
			continue
		case "n":
			break
		}

	}

	jsonData, err := json.MarshalIndent(feeds, "", "  ")
	if err != nil {
		err = fmt.Errorf("error while marshalling JSON: %v", err)
		fmt.Println(err)
	}
	err = os.WriteFile(jsonFileName, jsonData, 0644)
	if err != nil {
		err = fmt.Errorf("error while marshalling JSON: %v", err)
		fmt.Println(err)
	}

}
