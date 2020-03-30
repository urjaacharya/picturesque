package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
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
	if imageType == "favicon.ico" {
		imgFile, err := os.Create(filepath.Join(outputDir, imageType))
		buf := new(bytes.Buffer)
		err = png.Encode(buf, src)
		imgFile.Write(buf.Bytes())
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
		return
	}
	err = imaging.Save(src, filepath.Join(outputDir, imageType))

	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
