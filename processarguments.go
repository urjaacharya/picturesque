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

type arguments struct {
	Input_image      string
	Output           output
	Link             Link
	Site_webmanifest map[string]interface{}
	Icons            interface{}
}
type output struct {
	Images_path string
	HTML        html
}

type html struct {
	Path string
	Name string
}
type Link struct {
	Href_prefix string
	Href_suffix string
}
type Webmanifest struct {
	Background_color string
	Name             string
	Short_name       string
	Theme_color      string
}

// validateArguments: validate input arguments
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
	if args.Site_webmanifest["background_color"] == "" {
		args.Site_webmanifest["background_color"] = "#ffffff"
	}
	if args.Site_webmanifest["theme_color"] == "" {
		args.Site_webmanifest["theme_color"] = "#ffffff"
	}
	if args.Site_webmanifest["name"] == "" {
		args.Site_webmanifest["name"] = "default-name"
	}
	if args.Site_webmanifest["short_name"] == "" {
		args.Site_webmanifest["short_name"] = "default-short-name"
	}
}

//ReadArgs Reads user provided arguments from the input file
func ReadArgs() (string, string, map[string]interface{}, interface{}, Link, string) {
	inputArgsFile := flag.String("inputArgs", "", "REQUIRED: input arguments")
	flag.Usage = Usage
	flag.Parse()

	file, err := ioutil.ReadFile(*inputArgsFile)
	var args arguments
	err = json.Unmarshal(file, &args)
	validateArguments(args)

	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	iconsData := args.Icons.(map[string]interface{})
	hrefData := args.Link
	output := filepath.FromSlash(args.Output.Images_path)
	htmlFilepath := filepath.Join(args.Output.HTML.Path, args.Output.HTML.Name)

	return args.Input_image, output, args.Site_webmanifest, iconsData, hrefData, htmlFilepath
}
