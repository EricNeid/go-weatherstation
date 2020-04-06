all: test build

build:
	cd cmd/weatherstation && go build

run-windows:
	cd cmd/weatherstation && weatherstation.exe

run-linux:
	cd cmd/weatherstation && weatherstation

test:
	go test ./...