all: test build

build:
	cd cmd/weatherstation && go build

test:
	go test ./...

run-windows:
	cd cmd/weatherstation && weatherstation.exe

run-linux:
	cd cmd/weatherstation && weatherstation

cross-build-windows:
	cd cmd/weatherstation && go get github.com/lucor/fyne-cross/v2/cmd/fyne-cross && fyne-cross windows