all: test build

install-crosscompile-tools:
	go get github.com/lucor/fyne-cross/v2/cmd/fyne-cross

build:
	cd cmd/weatherstation && go build

run-windows:
	cd cmd/weatherstation && weatherstation.exe

run-linux:
	cd cmd/weatherstation && weatherstation

test:
	go test ./...