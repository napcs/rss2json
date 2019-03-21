package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
)

const AppVersion = "0.1.0"

func main() {
	version := flag.Bool("v", false, "Display the current version")
	prettyprint := flag.Bool("p", false, "Pretty-print JSON")

	// Custom message
	flag.Usage = func() {
		fmt.Printf("Usage: rss2json url\n")
		fmt.Printf("Returns JSON to STDOUT\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// version flag. print and exit
	if *version {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	// check if no options passed at all
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please specify a URL to an RSS feed")
		os.Exit(1)
	}

	// get the arguments
	url := flag.Args()[0]

	// convert
	convert(url, prettyprint)
}

func convert(url string, prettyprint *bool) {
	feed := fetchRSS(url)
	fmt.Println(convertToJson(feed, prettyprint))
}

func fetchRSS(url string) *gofeed.Feed {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		panic(err)
	}
	return feed
}

func convertToJson(feed *gofeed.Feed, prettyprint *bool) string {

	if *prettyprint {
		results, _ := json.MarshalIndent(feed, "", "  ")
		return string(results)
	} else {
		results, _ := json.Marshal(feed)
		return string(results)
	}
}
