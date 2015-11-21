package main

import "io/ioutil"
import "github.com/danzilio/go-yamlenc/Godeps/_workspace/src/gopkg.in/yaml.v2"

type Config struct {
	NodeList []string `yaml:"nodes"`
	Fail     bool     `yaml:"fail"`
}

func (c *Config) Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err == nil {
		yaml.Unmarshal([]byte(data), &c)
	}
	return err
}
