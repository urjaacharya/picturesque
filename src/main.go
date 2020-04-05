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
	"android-chrome-512x512.png": 512,
	"android-chrome-192x192.png": 192,
	"favicon-16x16.png":          16,
	"favicon-32x32.png":          32,
	"apple-touch-icon.png":       180,
	"favicon.ico":                48,
}

func main() {
	inputImage, outputDir, webManifestData, icons := ReadArgs()
	iconsData := icons.(map[string]interface{})
	imgFile, err := imaging.Open(inputImage)
	fmt.Println(iconsData)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	for key, value := range iconsData {
		generateFavicon(imgFile, key, value.(map[string]interface{}), outputDir)
	}
	outputManifest := generateWebManifest(faviconTypes, outputDir, webManifestData)
	generateIconsList(faviconTypes, outputDir, outputManifest)
}
