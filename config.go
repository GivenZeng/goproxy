package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

func parseConfig(conf interface{}, filePath string) error {
	if filePath == "" {
		filePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return err
		}
		filePath += "etc/conf.yaml"
	}
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return nil
}
