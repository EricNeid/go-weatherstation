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


linux-dependencies:
	sudo apt-get install libegl1-mesa-dev and xorg-dev


cross-build-windows:
	cd ../cmd/weatherstation && go get github.com/lucor/fyne-cross/v2/cmd/fyne-cross \
		&& fyne-cross windows \
		&& cp -r images fyne-cross/bin/windows-amd64 \
		&& echo replaceMe>fyne-cross/bin/windows-amd64/api.key

cross-build-raspberry:
	cd ../cmd/weatherstation && go get github.com/lucor/fyne-cross/v2/cmd/fyne-cross \
		&& fyne-cross linux -arch=arm \
		&& cp -r images fyne-cross/bin/linux-arm \
		&& echo replaceMe>fyne-cross/bin/linux-arm/api.key
