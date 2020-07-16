all: init resources test build distribute

init:
	go get fyne.io/fyne/cmd/fyne 

resources:
	fyne bundle -package res ./assets/ > res/bundle.go 

distribute:
	cd cmd/weatherstation \
	&& fyne package -icon ../../assets/app_icon.png

build:
	cd cmd/weatherstation && go build

test:
	go test ./...


linux-dependencies:
	sudo apt-get install libegl1-mesa-dev and xorg-dev


cross-build-windows:
	cd cmd/weatherstation && go get github.com/lucor/fyne-cross/v2/cmd/fyne-cross \
		&& fyne-cross windows \
		&& cp -r images fyne-cross/bin/windows-amd64 \
		&& echo replaceMe>fyne-cross/bin/windows-amd64/api.key

cross-build-raspberry:
	cd cmd/weatherstation && go get github.com/lucor/fyne-cross/v2/cmd/fyne-cross \
		&& fyne-cross linux -arch=arm \
		&& cp -r images fyne-cross/bin/linux-arm \
		&& echo replaceMe>fyne-cross/bin/linux-arm/api.key
