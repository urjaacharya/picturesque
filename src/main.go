package main

import (
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
	"android-chrome-512x512.png": 512,
	"android-chrome-192x192.png": 192,
	"favicon-16x16.png":          16,
	"favicon-32x32.png":          32,
	"apple-touch-icon.png":       180,
	"favicon.ico":                48,
}

func main() {
	inputImage, outputDir, webManifestData := ReadArgs()
	imgFile, err := imaging.Open(inputImage)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	for imageType, imageSize := range faviconTypes {
		generateFavicon(imgFile, imageType, imageSize, outputDir)
	}
	outputManifest := generateWebManifest(faviconTypes, outputDir, webManifestData)
	generateIconsList(faviconTypes, outputDir, outputManifest)
}
