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
