package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	config map[string]string
}

func readConfigFromJson(filename string) (*Config, error) {
	contents, err = ioutil.ReadFile(filename)
	config := Config{make(map[string]string)}
	if err != nil {
		return nil, err
	}
	data, err = json.Decode(string(contents))
	if err != nil {
		return nil, err
	}
	for key, value := range data {
		config[key] = value
	}
	return &config, nil
}
func main() {
	readConfigFromFileJson("test.json")
}
