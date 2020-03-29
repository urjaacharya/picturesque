package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

func generateFavicon(inputImage image.Image, imageType string, imgSize int, outputDir string) {
	src := imaging.Resize(inputImage, imgSize, imgSize, imaging.Lanczos)
	fmt.Println(outputDir)
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}
	resizeErr := imaging.Save(src, filepath.Join(outputDir, imageType+".png"))
	if resizeErr != nil {
		log.Fatalf("failed to resize image: %v", resizeErr)
	}
}
