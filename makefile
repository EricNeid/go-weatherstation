build:
	cd cmd/weatherstation && go build

run:
	cd cmd/weatherstation && weatherstation.exe

test:
	go test ./...