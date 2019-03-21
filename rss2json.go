package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
)

const appVersion = "0.1.0"

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
		fmt.Println(appVersion)
		os.Exit(0)
	}

	// check if no options passed at all
	if len(flag.Args()) == 0 {
		fmt.Println("Please specify a URL to an RSS feed")
		os.Exit(1)
	}

	// get the arguments
	url := flag.Args()[0]

	// convert
	if err := convert(url, *prettyprint); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func convert(url string, prettyprint bool) error {
	feed, err := fetchRSS(url)
	if err != nil {
		return err
	}
	result, err := convertToJson(feed, prettyprint)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func fetchRSS(url string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	return feed, err
}

func convertToJson(feed *gofeed.Feed, prettyprint bool) (string, error) {

	if prettyprint {
		results, err := json.MarshalIndent(feed, "", "  ")
		return string(results), err
	}

	results, err := json.Marshal(feed)
	return string(results), err
}
