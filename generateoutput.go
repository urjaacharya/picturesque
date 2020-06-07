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

// AddIconsListToWebManifest Adds icons' list to site webmanifest file
func AddIconsListToWebManifest(icons map[string]interface{}, outputDir string, outputData map[string]interface{}) {
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

// GenerateWebManifest Generates site webmanifest file
func GenerateWebManifest(outputDir string, webManifestData map[string]interface{}) map[string]interface{} {
	outputData := make(map[string]interface{})

	for key, value := range webManifestData {
		outputData[key] = value
		fmt.Println(outputData[key])
	}

	//give default values for some key value pairs
	outputData["name"] = func() string {
		if webManifestData["name"] != "" {
			return fmt.Sprintf("%v", webManifestData["name"])
		}
		return ""
	}()
	outputData["short_name"] = func() string {
		if webManifestData["short_name"] != "" {
			return fmt.Sprintf("%v", webManifestData["short_name"])
		}
		return ""
	}()
	outputData["background_color"] = func() string {
		if webManifestData["background_color"] != "" {
			return fmt.Sprintf("%v", webManifestData["background_color"])
		}
		return "#ffffff"
	}()
	outputData["theme_color"] = func() string {
		if webManifestData["theme_color"] != "" {
			return fmt.Sprintf("%v", webManifestData["theme_color"])
		}
		return "#ffffff"
	}()

	jsonString, _ := json.Marshal(outputData)
	err := ioutil.WriteFile(filepath.Join(outputDir, "site.webmanifest"), jsonString, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}

	return outputData
}

// GenerateFavicon Generates favicon of specified size
func GenerateFavicon(inputImage image.Image, imageName string, imgData map[string]interface{}, outputDir string) {
	//to do: crop image to fix aspect ratio
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
		defer imgFile.Close()

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

// GenerateHTML Generates HTML file with favicon definitions
func GenerateHTML(icons map[string]interface{}, hrefData Link, filePath string) {
	file, err := os.Create(filePath + ".html")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
	prefix := hrefData.Href_prefix
	suffix := hrefData.Href_suffix

	for imageName, imageData := range icons {
		data := imageData.(map[string]interface{})
		rel := data["rel"].([]interface{})
		height := fmt.Sprintf("%v", data["height"])
		width := fmt.Sprintf("%v", data["width"])
		for item := range rel {
			file.WriteString("<link rel=" + `"` + rel[item].(string) + `"` + " href=" + `"` + prefix + imageName + ".png" + suffix + `"` + " sizes=" + `"` + width + "x" + height + `"` + " type=" + `"` + "image/png" + `"` + "/>\n")
		}
	}
}
