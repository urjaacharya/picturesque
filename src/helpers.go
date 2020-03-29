package main

import (
	"flag"
	"fmt"
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

//ReadArgs Reads user provided arguments
func ReadArgs() (string, string) {
	inputImage := flag.String("input", "", "REQUIRED: input image")
	outputDir := flag.String("outputDir", filepath.Join("output", "favicons"), "OPTIONAL: output folder to store the generated favicons")

	flag.Usage = Usage
	flag.Parse()

	if *inputImage == "" {
		fmt.Println("ERROR: Input image not provided")
		os.Exit(1)
	}

	output := filepath.FromSlash(*outputDir)
	return *inputImage, output
}
