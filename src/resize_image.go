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

func generateIconsList(icons map[string]interface{}, outputDir string, outputData map[string]interface{}) {
	var iconsList []map[string]string

	for imageName, imageData := range icons {
		imageHeight := imageData.(map[string]interface{})["height"]
		imageWidth := imageData.(map[string]interface{})["width"]
		if imageName != "favicon.ico" {
			data := make(map[string]string)
			data["src"] = filepath.Join(outputDir, imageName+".png")
			data["sizes"] = fmt.Sprintf("%v", imageWidth) + "x" + fmt.Sprintf("%v", imageHeight)
			data["type"] = "image/png"
			iconsList = append(iconsList, data)
		}
	}

	outputData["icons"] = iconsList
	jsonString, _ := json.Marshal(outputData)
	err := ioutil.WriteFile(filepath.Join(outputDir, "site.webmanifest"), jsonString, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
}

func generateWebManifest(outputDir string, webManifestData interface{}) map[string]interface{} {
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

func generateFavicon(inputImage image.Image, imageName string, imgData map[string]interface{}, outputDir string) {
	src := imaging.Resize(inputImage, int(imgData["width"].(float64)), int(imgData["height"].(float64)), imaging.Lanczos)
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}
	if imageName == "favicon.ico" {
		imgFile, err := os.Create(filepath.Join(outputDir, imageName+".png"))
		buf := new(bytes.Buffer)
		err = png.Encode(buf, src)
		imgFile.Write(buf.Bytes())
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
		return
	}
	err = imaging.Save(src, filepath.Join(outputDir, imageName+".png"))

	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

}
