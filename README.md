# rss2json

Command line utility to convert RSS feeds to JSON.

## Usage

Pass the URL to your RSS feed and save the result to a file using shell redirection.

```
$ rss2json https://example.com/rss/ > out.json
```

Use the `-p` switch to pretty-print the output.


## Installation

To install, download the latest release to your system and copy the executable to a location on your path. 

Alternatively, clone the repository and build the version you need from source with

```
$ make windows mac linux
```

## Contributing

Please contribute. Clone this repo. Install the go-feed library:

```
$ go get github.com/mmcdole/gofeed
```

Then make your changes. Build the packages with 

```
$ make windows mac linux
```

Test them on those operating systems, and then send in a PR.

## History

* 2019-03-20 - v0.1.0
  * initial release

## License

Apache 2. See LICENSE file.
