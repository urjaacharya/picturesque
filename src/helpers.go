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

// validate arguments: validate input arguments
func validateArguments(args arguments) {
	if args.Input_image == "" {
		log.Fatalf("input_image missing in input json")
	}
	if args.Output.Images_path == "" {
		log.Fatalf("Path to output directory for images missing in input json")
	}
	if args.Output.HTML.Path == "" {
		log.Fatalf("Path to output directory for html file missing in input json")
	}
	if args.Output.HTML.Name == "" {
		log.Fatalf("Name of HTML file is missing in input json")
	}
	if args.Site_webmanifest.Background_color == "" {
		args.Site_webmanifest.Background_color = "#ffffff"
	}
	if args.Site_webmanifest.Theme_color == "" {
		args.Site_webmanifest.Theme_color = "#ffffff"
	}
	if args.Site_webmanifest.Name == "" {
		args.Site_webmanifest.Name = "default-name"
	}
	if args.Site_webmanifest.Short_name == "" {
		args.Site_webmanifest.Short_name = "default-short-name"
	}
}

//ReadArgs Reads user provided arguments
func ReadArgs() (string, string, interface{}, interface{}, map[string]interface{}, string) {
	inputArgsFile := flag.String("inputArgs", "", "REQUIRED: input arguments")
	flag.Usage = Usage
	flag.Parse()
	var names arguments

	file, err := ioutil.ReadFile(*inputArgsFile)
	var args map[string]interface{}
	err = json.Unmarshal(file, &args)
	err = json.Unmarshal(file, &names)
	fmt.Println(names)
	validateArguments(names)
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
	output := filepath.FromSlash(fmt.Sprintf("%v", names.Output.Images_path))
	html := imagesOutputDir["html"].(map[string]interface{})
	htmlFilepath := filepath.Join(html["path"].(string), html["name"].(string))
	return fmt.Sprintf("%v", names.Input_image), output, args["site_webmanifest"], iconsData, hrefData, htmlFilepath
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
