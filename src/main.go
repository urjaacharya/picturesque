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
	"androidChrome512": 512,
	"androidChrome192": 192,
	"favicon16":        16,
	"favicon32":        32,
	"appleTouchIcon":   180,
}

func main() {
	imgFile, err := imaging.Open("data/beach-soft-light.jpg")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	for imageType, imageSize := range faviconTypes {
		generateFavicon(imgFile, imageType, imageSize)
	}
	fmt.Println(androidChrome192)
}
