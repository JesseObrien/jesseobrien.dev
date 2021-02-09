.PHONY: install
install:
	go get github.com/rakyll/statik
	go get -u github.com/cosmtrek/air

.PHONY: build
build:
	statik -f src=public
	go build -ldflags="-s -w" -o ./dist/jesseobrien main.go

.PHONY: run
run:
	air
