DIR := ${CURDIR}


.PHONY: init
init:
	go install fyne.io/fyne/v2/cmd/fyne@v2.2.1


.PHONY: resources
resources: init
	fyne bundle -package assets ./resources/ > ./assets/bundle.go 


.PHONY: build
build: init resources
	cd cmd/weatherstation && fyne package -icon ../../resources/app_icon.png


.PHONY: test
test:
	go test ./...


.PHONY: lint
lint:
	docker run -it --rm \
		-e CGO_ENABLED=0 \
		-w /app -v ${DIR}:/app \
		golangci/golangci-lint:v1.52.2 \
		golangci-lint run ./...


.PHONY: cross-build-windows
cross-build-windows:
	cd ../cmd/weatherstation && go get github.com/lucor/fyne-cross/v2/cmd/fyne-cross \
		&& fyne-cross windows \
		&& cp -r images fyne-cross/bin/windows-amd64 \
		&& echo replaceMe>fyne-cross/bin/windows-amd64/api.key


.PHONY: cross-build-raspberry
cross-build-raspberry:
	cd ../cmd/weatherstation && go get github.com/lucor/fyne-cross/v2/cmd/fyne-cross \
		&& fyne-cross linux -arch=arm \
		&& cp -r images fyne-cross/bin/linux-arm \
		&& echo replaceMe>fyne-cross/bin/linux-arm/api.key
