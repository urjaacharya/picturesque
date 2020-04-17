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

func generateWebManifest(outputDir string, webManifestData Webmanifest) map[string]interface{} {
	//inputData := webManifestData.(map[string]interface{})
	outputData := make(map[string]interface{})
	outputData["name"] = func() string {
		if webManifestData.Name != "" {
			return fmt.Sprintf("%v", webManifestData.Name)
		}
		return ""
	}()
	outputData["short_name"] = func() string {
		if webManifestData.Short_name != "" {
			return fmt.Sprintf("%v", webManifestData.Short_name)
		}
		return ""
	}()
	outputData["background_color"] = func() string {
		if webManifestData.Background_color != "" {
			return fmt.Sprintf("%v", webManifestData.Background_color)
		}
		return "#ffffff"
	}()
	outputData["theme_color"] = func() string {
		if webManifestData.Theme_color != "" {
			return fmt.Sprintf("%v", webManifestData.Theme_color)
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

func generateHTML(icons map[string]interface{}, hrefData Link, filePath string) {
	file, err := os.Create(filePath + ".html") // Truncates if file already exists, be careful!
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close() // Make sure to close the file when you're done
	prefix := hrefData.Href_prefix
	suffix := hrefData.Href_suffix

	for imageName, imageData := range icons {
		data := imageData.(map[string]interface{})
		rel := data["rel"].([]interface{})
		height := fmt.Sprintf("%v", data["height"])
		width := fmt.Sprintf("%v", data["width"])
		for item := range rel {
			file.WriteString("<link rel=" + `"` + rel[item].(string) + `"` + " href=" + `"` + prefix + imageName + ".png" + suffix + `"` + " sizes=" + `"` + width + "x" + height + `"` + " type=" + `"` + "image/png" + `"` + "/>")
		}

	}
}
