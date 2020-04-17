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
	Site_webmanifest Webmanifest
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

// type icons struct {
// 	Favicon16  faviconInfo `json:"favicon-16"`
// 	Favicon32  faviconInfo `json:"favicon-32"`
// 	Favicon120 faviconInfo `json:"favicon-120"`
// 	Favicon128 faviconInfo `json:"favicon-128"`
// 	Favicon152 faviconInfo `json:"favicon-152"`
// 	Favicon167 faviconInfo `json:"favicon-167"`
// 	Favicon180 faviconInfo `json:"favicon-180"`
// 	Favicon192 faviconInfo `json:"favicon-192"`
// 	Favicon196 faviconInfo `json:"favicon-196"`
// }

// type faviconInfo struct {
// 	Width  uint16
// 	Height uint16
// 	Rel    []string
// }

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
func ReadArgs() (string, string, Webmanifest, interface{}, Link, string) {
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
