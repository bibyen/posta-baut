package main

import (
	"log"
	"os"

	"github.com/bibyen/posta-baut/cmd/client"
	config "github.com/bibyen/posta-baut/cmd/config"
	"github.com/bibyen/posta-baut/cmd/svc"

	"github.com/bibyen/posta-baut/cmd/server"
	yaml "github.com/goccy/go-yaml"
)

func loadConfig(path string) (config.AppConfig, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// 3. Declare a variable of your struct type to hold the data
	var config config.AppConfig

	// 4. Unmarshal the byte slice into the struct
	if err := yaml.Unmarshal(fileBytes, &config); err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}
	return config, nil
}

func main() {
	cfg, err := loadConfig("./config/local.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
		return
	}

	// Create teams service
	cli, err := client.NewClient(cfg.SenderConfig, cfg.LookupClientConfig)
	if err != nil {
		log.Fatalf("Failed to create client: %v\n", err)
		return
	}
	svc := svc.NewTeamsService(&cli)

	// Create and run the server
	server, err := server.NewServer(svc)
	if err != nil {
		log.Fatalf("Failed to create server: %v\n", err)
		return
	}
	err = server.Run()
	if err != nil {
		log.Fatalf("error running server: %v\n", err)
	}
}
