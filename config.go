package main

import "io/ioutil"
import "github.com/danzilio/go-yamlenc/Godeps/_workspace/src/gopkg.in/yaml.v2"

type Config struct {
	NodeList []string `yaml:"nodes"`
	Fail     bool     `yaml:"fail"`
}

type StringNode struct {
	NodeList string `yaml:"nodes"`
	Fail     bool   `yaml:"fail"`
}

func (c *Config) Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err == nil {
		err = yaml.Unmarshal([]byte(data), &c)
		if err != nil {
			string_node := StringNode{}
			err = yaml.Unmarshal([]byte(data), &string_node)
			if err == nil {
				c.NodeList = append(c.NodeList, string_node.NodeList)
				c.Fail = string_node.Fail
			}
		}
	}
	return err
}
