package main

import "io/ioutil"
import "github.com/danzilio/go-yamlenc/Godeps/_workspace/src/gopkg.in/yaml.v2"

type Config struct {
	NodeList NodeList `yaml:"nodes"`
	Fail     bool     `yaml:"fail"`
}

type NodeList []string

func (n *NodeList) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	err := unmarshal(&str)
	if err == nil {
		*n = []string{str}
		return nil
	}

	var slice []string
	err = unmarshal(&slice)
	if err == nil {
		*n = slice
		return nil
	}

	return err
}

func (c *Config) Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err == nil {
		err = yaml.Unmarshal([]byte(data), c)
	}
	return err
}
