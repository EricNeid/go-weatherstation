DIR := ${CURDIR}


.PHONY: init
init:
	go get fyne.io/fyne/v2/cmd/fyne@v2.2.1


.PHONY: resources
resources: init
	fyne bundle -package assets ./assets/ > internal/assets/bundle.go 


.PHONY: build
build: init
	fyne package -icon assets/app_icon.png -sourceDir cmd/weatherstation/


.PHONY: test
test:
	go test ./...


.PHONY: lint
lint:
	docker run -it --rm \
		-e CGO_ENABLED=0 \
		-w /app -v ${DIR}:/app \
		golangci/golangci-lint:v1.50.1 \
		golangci-lint run ./...
