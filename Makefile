.PHONY: test

all: init resources test build

init:
	go get fyne.io/fyne/v2/cmd/fyne 

resources:
	fyne bundle -package assets ./assets/ > internal/assets/bundle.go 

build:
	cd cmd/weatherstation && go build && fyne package -icon ../../assets/app_icon.png

test:
	go test ./...
