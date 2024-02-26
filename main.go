package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	configFile string
	port       string
)

type Config struct {
	MailcowHost     string `yaml:"MAILCOW_HOSTNAME"`
	MailcowToken    string `yaml:"MAILCOW_TOKEN"`
	PrometheusToken string `yaml:"PROMETHEUS_TOKEN"`
}

func parseFlags() {
	flag.StringVar(&configFile, "config", "./config.yaml", "Location of the config.yaml file that should be used")
	flag.StringVar(&port, "port", "9099", "The port that the exporter should listen to")

	flag.Parse()
}

// Read and deserialize the config file, return a pointer to the result (type Config)
func readConfig() (*Config, error) {
	fileContent, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("reading %v: %w", configFile, err)
	}

	var result Config
	err = yaml.Unmarshal(fileContent, &result)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling YAML data: %w", err)
	}

	return &result, nil
}

func main() {
	parseFlags()

	conf, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf)
}
