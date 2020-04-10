package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//Usage Displays usage of defined arguments
func Usage() {
	fmt.Println("\npicturesque version 0.0 'hang tight' USAGE")
	fmt.Println("========================================")
	fmt.Println("-image   REQUIRED: input image.")
	fmt.Println("-outputDir   REQUIRED: directory to store the output favicons")
	os.Exit(1)
}

func verifyArguments() {
	inputArgsFile := flag.String("inputArgsFile", "", "REQUIRED: input json with input arguments")
	flag.Usage = Usage
	flag.Parse()
	file, err := ioutil.ReadFile(*inputArgsFile)
	var args map[string]interface{}
	err = json.Unmarshal(file, &args)

	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Println(args["new"])
	if args["input_image"] == nil {
		log.Fatalf("input_image missing in input json")
	}
	if args["output"] == nil {
		log.Fatalf("Path to output directory missing in input json")
	}
	if args["output"] == nil {
		log.Fatalf("Path to output directory missing in input json")
	}
}

type arguments struct {
	Input_image      string
	Output           output
	Link             link
	Site_webmanifest webmanifest
	Icons            icons
}
type output struct {
	Images_path string
	HTML        html
}

type html struct {
	Path string
	Name string
}
type link struct {
	Href_prefix string
	Href_suffix string
}
type webmanifest struct {
	Background_color string
	Name             string
	Short_name       string
	Theme_color      string
}

type icons struct {
	Favicon16  faviconInfo `json:"favicon-16"`
	Favicon32  faviconInfo `json:"favicon-32"`
	Favicon120 faviconInfo `json:"favicon-120"`
	Favicon128 faviconInfo `json:"favicon-128"`
	Favicon152 faviconInfo `json:"favicon-152"`
	Favicon167 faviconInfo `json:"favicon-167"`
	Favicon180 faviconInfo `json:"favicon-180"`
	Favicon192 faviconInfo `json:"favicon-192"`
	Favicon196 faviconInfo `json:"favicon-196"`
}

type faviconInfo struct {
	Width  uint16
	Height uint16
	Rel    []string
}

func test() {
	new_file, err := ioutil.ReadFile("./test.json")
	var names arguments
	err = json.Unmarshal(new_file, &names)
	fmt.Println(names)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}

//ReadArgs Reads user provided arguments
func ReadArgs() (string, string, interface{}, interface{}, map[string]interface{}, string) {
	test()
	inputArgsFile := flag.String("inputArgs", "", "REQUIRED: input arguments")
	flag.Usage = Usage
	flag.Parse()
	var names arguments

	file, err := ioutil.ReadFile(*inputArgsFile)
	var args map[string]interface{}
	err = json.Unmarshal(file, &args)
	err = json.Unmarshal(file, &names)
	fmt.Println(names)
	// To do: add function to check input arguments in input json
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Println(args["new"])
	if args["input_image"] == nil {
		log.Fatalf("input_image missing in input json")
	}
	if args["output"] == nil {
		log.Fatalf("Path to output directory missing in input json")
	}

	imagesOutputDir := args["output"].(map[string]interface{})
	iconsData := args["icons"].(map[string]interface{})
	hrefData := args["link"].(map[string]interface{})
	output := filepath.FromSlash(fmt.Sprintf("%v", imagesOutputDir["images_path"]))
	html := imagesOutputDir["html"].(map[string]interface{})
	htmlFilepath := filepath.Join(html["path"].(string), html["name"].(string))
	return fmt.Sprintf("%v", args["input_image"]), output, args["site_webmanifest"], iconsData, hrefData, htmlFilepath
}

func generateHTML(icons map[string]interface{}, hrefData map[string]interface{}, filePath string) {
	file, err := os.Create(filePath + ".html") // Truncates if file already exists, be careful!
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close() // Make sure to close the file when you're done
	prefix := hrefData["href_prefix"].(string)
	suffix := hrefData["href_suffix"].(string)

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
