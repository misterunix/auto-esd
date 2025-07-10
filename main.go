package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

//go:embed smallimage.tmpl
var smallImageTemplate string

//go:embed largeimage.tmpl
var largeImageTemplate string

//go:embed samplers.txt
var samplers string

func main() {

	var createConfig bool
	var loadCongig bool
	var configFile string
	var modelsDir string
	flag.BoolVar(&createConfig, "cc", false, "Create a new configuration file")
	flag.StringVar(&modelsDir, "md", "", "Path to the models directory")
	flag.StringVar(&configFile, "c", "config.json", "Path to the configuration file")
	flag.BoolVar(&loadCongig, "lc", false, "Load the configuration file")
	flag.Parse()

	sd := newStable()

	//
	//
	//

	if createConfig {
		// Create a new configuration file
		fmt.Println("Creating a new configuration file at", configFile)

		if modelsDir == "" {
			fmt.Fprintln(os.Stderr, "modelsDir is required when creating a new configuration file")
			os.Exit(1)
		}

		if _, err := os.Stat(modelsDir); os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, "modelsDir does not exist:", modelsDir)
			os.Exit(1)
		}

		sd.ModelsDir = modelsDir // Set the models directory from the command line argument

		jsonData, err := json.MarshalIndent(sd, "", "  ")
		if err != nil {
			fmt.Println("Error creating configuration file:", err)
			os.Exit(0)
		}
		os.WriteFile(configFile, jsonData, 0644)
		fmt.Println("Configuration file created successfully.")
		os.Exit(0)
	} else if loadCongig {
		// Load the configuration file
		data, err := os.ReadFile(configFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ReadFile:", err)
			os.Exit(1)
		}

		err = json.Unmarshal(data, &sd)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unmarshal:", err)
			os.Exit(1)
		}
	}

	//
	//
	//

	if sd.ModelsDir == "" {
		fmt.Fprintln(os.Stderr, "ModelsDir is required when creating a new configuration file")
		os.Exit(1)
	}

	if _, err := os.Stat(sd.ModelsDir); os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "ModelsDir does not exist:", sd.ModelsDir)
		os.Exit(1)
	}

	//
	//
	//

	LoadModels(sd) // Load the models from the directory

	//
	//
	//

	//
	//
	//

	err := sit(sd) // Generate the small image script and run it
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error generating small image script:", err)
		os.Exit(1)
	}

	err = lit(sd) // Generate the large image script and run it
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error generating large image script:", err)
		os.Exit(1)
	}

}

func newStable() *stable {
	return &stable{
		Seed:             0,
		Width:            512,
		Height:           512,
		Prompt:           "",
		Nprompt:          "",
		Model:            "",
		Sampler:          "",
		Steps:            50,
		DateTime:         "",
		SmallImage:       "",
		LargeImage:       "",
		SmallImagePython: "smallImage.py",
		LargeImagePython: "largeImage.py",
	}
}

// get all the models from the directory
func LoadModels(sd *stable) error {
	fmt.Println("Loading Models")

	// get a list of models from a directory
	dir, err := os.ReadDir(sd.ModelsDir)
	if err != nil {
		return err
	}

	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		if strings.HasSuffix(file.Name(), ".txt") {
			continue
		}
		if strings.HasSuffix(file.Name(), ".safetensors") {
			fmt.Println("Found model: ", file.Name())
			sd.Models = append(sd.Models, file.Name())
		}
	}

	if len(sd.Models) == 0 {
		return fmt.Errorf("no models found")
	}

	fmt.Println("Models loaded successfully:", len(sd.Models))

	return nil
}
