package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

func generateIconsList(favicons map[string]int, outputDir string, outputData map[string]interface{}) {
	iconsList := make([]map[string]string, 0)
	data := make(map[string]string)
	for imageType, imageSize := range faviconTypes {
		if imageType != "favicon.ico" {
			data["src"] = filepath.Join(outputDir, imageType)
			data["sizes"] = fmt.Sprintf("%v", imageSize) + "x" + fmt.Sprintf("%v", imageSize)
			data["type"] = "image/png"
			iconsList = append(iconsList, data)
			fmt.Println(iconsList)
		}
	}
	outputData["icons"] = iconsList
	jsonString, _ := json.Marshal(outputData)
	err := ioutil.WriteFile(filepath.Join(outputDir, "site.webmanifest"), jsonString, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
}

func generateWebManifest(favicons map[string]int, outputDir string, webManifestData interface{}) map[string]interface{} {
	inputData := webManifestData.(map[string]interface{})
	outputData := make(map[string]interface{})
	outputData["name"] = func() string {
		if inputData["name"] != nil {
			return fmt.Sprintf("%v", inputData["name"])
		}
		return ""
	}()
	outputData["short_name"] = func() string {
		if inputData["short_name"] != nil {
			return fmt.Sprintf("%v", inputData["short_name"])
		}
		return ""
	}()
	outputData["background_color"] = func() string {
		if inputData["background_color"] != nil {
			return fmt.Sprintf("%v", inputData["background_color"])
		}
		return "#ffffff"
	}()
	outputData["theme_color"] = func() string {
		if inputData["theme_color"] != nil {
			return fmt.Sprintf("%v", inputData["theme_color"])
		}
		return "#ffffff"
	}()

	jsonString, _ := json.Marshal(outputData)
	err := ioutil.WriteFile(filepath.Join(outputDir, "site.webmanifest"), jsonString, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	fmt.Println(inputData["name"])
	return outputData
}
func generateFavicon(inputImage image.Image, imageType string, imgSize int, outputDir string) {
	src := imaging.Resize(inputImage, imgSize, imgSize, imaging.Lanczos)
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
