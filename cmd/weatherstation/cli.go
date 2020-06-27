package main

import (
	"flag"
	"fmt"
	"os"
)

type args struct {
	fullscreen bool
	imageDir   string
	keyFile    string
}

func parseArgs() args {
	flag.Usage = func() {
		fmt.Printf("Usage: %s\n", os.Args[0])
		flag.PrintDefaults()
	}

	var args args
	flag.BoolVar(&args.fullscreen, "fullscreen", false, "show app in fullscreen mode")
	flag.StringVar(&args.imageDir, "images", "images", "directory, containing images for the screensave")
	flag.StringVar(&args.keyFile, "key", "api.key", "file, containing api key for openweather")

	flag.Parse()

	return args
}
