package main

import (
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	inputImage, outputDir, webManifestData, icons, hrefData, htmlFilepath := ReadArgs()
	iconsData := icons.(map[string]interface{})
	imgFile, err := imaging.Open(inputImage)

	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// generates favicons for all specified sizes
	for key, value := range iconsData {
		GenerateFavicon(imgFile, key, value.(map[string]interface{}), outputDir)
	}
	outputManifest := GenerateWebManifest(outputDir, webManifestData)
	AddIconsListToWebManifest(iconsData, outputDir, outputManifest)
	GenerateHTML(iconsData, hrefData, htmlFilepath)
}
