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

	for key, value := range iconsData {
		generateFavicon(imgFile, key, value.(map[string]interface{}), outputDir)
	}
	outputManifest := generateWebManifest(outputDir, webManifestData)
	generateIconsList(iconsData, outputDir, outputManifest)
	generateHTML(iconsData, hrefData, htmlFilepath)
}
