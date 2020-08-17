.PHONY: test

all: init resources test build distribute

init:
	go get fyne.io/fyne/cmd/fyne 

resources:
	fyne bundle -package assets ./assets/ > internal/assets/bundle.go 

distribute:
	cd cmd/weatherstation \
	&& fyne package -icon ../../assets/app_icon.png

build:
	cd cmd/weatherstation && go build

test:
	go test ./...
