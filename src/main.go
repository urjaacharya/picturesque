package main

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
)

const (
	androidChrome512 = "androidChrome512"
	androidChrome192 = "androidChrome192"
	favicon16        = "favicon16"
	favicon32        = "favicon32"
	appleTouchIcon   = "appleTouchIcon"
)

var faviconTypes = map[string]int{
	"android-chrome-512x512": 512,
	"android-chrome-192x192": 192,
	"favicon-16x16":          16,
	"favicon-32x32":          32,
	"apple-touch-icon":       180,
}

func main() {
	inputImage, outputDir := ReadArgs()
	fmt.Println(inputImage, outputDir)
	imgFile, err := imaging.Open(inputImage)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	for imageType, imageSize := range faviconTypes {
		generateFavicon(imgFile, imageType, imageSize, outputDir)
	}
	fmt.Println(androidChrome192)
}
