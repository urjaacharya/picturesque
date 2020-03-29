package main

import (
	"image"
	"log"

	"github.com/disintegration/imaging"
)

func generateFavicon(inputImage image.Image, imageType string, imgSize int) {
	src := imaging.Resize(inputImage, imgSize, imgSize, imaging.Lanczos)
	resizeErr := imaging.Save(src, imageType+".png")
	if resizeErr != nil {
		log.Fatalf("failed to resize image: %v", resizeErr)
	}
}
