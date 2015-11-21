package main

import "strings"
import "io/ioutil"
import "gopkg.in/yaml.v2"

type EncNode struct {
	Role        string            `yaml:"role"`
	Environment string            `yaml:"environment"`
	Classes     []string          `yaml:"classes"`
	Parameters  map[string]string `yaml:"parameters"`
}

type PuppetNode struct {
	Environment string
	Classes     []string
	Parameters  map[string]string
}

func Nodes(file string) map[string]EncNode {
	nodes := make(map[string]EncNode)
	data, err := ioutil.ReadFile(file)
	if err == nil {
		yaml.Unmarshal([]byte(data), &nodes)
	}
	return nodes
}

func (n *EncNode) ToPuppetNode() PuppetNode {
	node := PuppetNode{}

	if len(n.Parameters) > 0 {
		node.Parameters = n.Parameters
	} else {
		node.Parameters = make(map[string]string)
	}

	if len(n.Classes) > 0 {
		node.Classes = n.Classes
	}

	if len(n.Role) > 0 {
		node.Classes = append(n.Classes, n.Role)
		role := strings.Replace(n.Role, "roles::", "", 1)
		role = strings.Replace(role, "::", "/", -1)
		node.Parameters["role"] = role
	}

	node.Environment = n.Environment
	return node
}
