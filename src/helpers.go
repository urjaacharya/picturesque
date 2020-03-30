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

//ReadArgs Reads user provided arguments
func ReadArgs() (string, string, interface{}) {
	inputArgsFile := flag.String("inputArgs", "", "REQUIRED: input arguments")
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
	if args["output_dir"] == nil {
		log.Fatalf("Path to output directory missing in input json")
	}

	output := filepath.FromSlash(fmt.Sprintf("%v", args["output_dir"]))
	return fmt.Sprintf("%v", args["input_image"]), output, args["site_webmanifest"]
}
