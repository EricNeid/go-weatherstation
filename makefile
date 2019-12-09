build:
	cd cmd/weatherstation && go build

run:
	cmd/weatherstation/weatherstation.exe

test:
	go test ./...