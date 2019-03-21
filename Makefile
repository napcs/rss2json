version = $(go run rss2json -v)

mac:
	mkdir -p dist/mac
	cp README.md dist/mac/README
	cp LICENSE dist/mac/LICENSE
	env GOOS=darwin GOARCH=amd64 go build -o dist/mac/rss2json
	cd dist/mac; zip -9 rss2json-macos-${version}.zip rss2json README LICENSE

linux:
	mkdir -p dist/linux
	cp README.md dist/linux/README
	cp LICENSE dist/linux/LICENSE
	env GOOS=linux GOARCH=amd64 go build -o dist/linux/rss2json
	cd dist/linux; zip -9 rss2json-linux-${version}.zip rss2json README LICENSE

windows:
	mkdir -p dist/windows
	cp README.md dist/windows/README
	cp LICENSE dist/windows/LICENSE
	env GOOS=windows GOARCH=amd64 go build -o dist/windows/rss2json.exe
	cd dist/windows; zip -9 rss2json-windows-${version}.zip rss2json.exe README LICENSE

clean:
	rm -rf dist/
