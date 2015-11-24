package main

import "os"
import "fmt"
import "io/ioutil"
import "regexp"
import "errors"
import "github.com/danzilio/go-yamlenc/Godeps/_workspace/src/github.com/codegangsta/cli"

func main() {
	app := cli.NewApp()
	app.Name = "yamlenc"
	app.Usage = "A very simple external node classifier (ENC) for Puppet"
	app.Version = Version

	var fail bool = false
	var config_file string
	config := Config{}
	cli_nodes := cli.StringSlice{}

	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "nodes, n",
			Usage: "Path to a data file to add to the end of the 'nodes' array.",
			Value: &cli_nodes,
		},
		cli.StringFlag{
			Name:        "config, c",
			Value:       "/etc/puppet/enc.yaml",
			Usage:       "Path to configuration file (Default: /etc/puppet/enc.yaml).",
			Destination: &config_file,
		},
		cli.BoolFlag{
			Name:        "fail, f",
			Usage:       "Fail if no nodes are found.",
			Destination: &fail,
		},
	}

	app.Action = func(c *cli.Context) {
		config.Load(config_file)
		if fail {
			config.Fail = fail
		}

		if len(c.Args()) == 0 {
			fmt.Println("ERROR: Didn't specify a node name to look up!")
			os.Exit(1)
		}

		for _, element := range cli_nodes {
			config.NodeList = append(config.NodeList, element)
		}

		node_list := CollectNodes(config.NodeList)

		if len(config.NodeList) == 0 {
			fmt.Println("No node files specified")
			os.Exit(1)
		}

		node, error := Lookup(c.Args()[0], node_list)

		if error == nil {
			puppet_node := node.ToPuppetNode()
			fmt.Println(puppet_node.String())
		} else {
			if fail == true {
				fmt.Println("No node found.")
				os.Exit(1)
			} else {
				fmt.Println("{}")
			}
		}
	}

	app.Run(os.Args)
}

func Dir(dir string, regex string) []string {
	var file_collection []string

	files, error := ioutil.ReadDir(dir)

	if error != nil {
		panic(error)
	}

	for _, file := range files {
		if file.IsDir() == true {
			dir_name := fmt.Sprintf("%s/%s", dir, file.Name())
			found_files := Dir(dir_name, regex)
			for _, f := range found_files {
				file_collection = append(file_collection, f)
			}
		} else {
			match, _ := regexp.MatchString(regex, file.Name())
			if match == true {
				file_collection = append(file_collection, fmt.Sprintf("%s/%s", dir, file.Name()))
			}
		}
	}

	return file_collection
}

func CollectNodes(nodes []string) []string {
	var collection []string

	for _, node := range nodes {
		info, err := os.Stat(node)
		if err == nil && info.IsDir() == true {
			for _, file := range Dir(node, "\\.yaml$") {
				collection = append(collection, file)
			}
		} else {
			collection = append(collection, node)
		}
	}

	return collection
}

func Lookup(name string, nodes []string) (EncNode, error) {
	for _, node := range nodes {
		found, error := search(name, Nodes(node))
		if error == nil {
			return found, nil
		}
	}
	return EncNode{}, errors.New("No node found!")
}

func search(name string, nodes map[string]EncNode) (EncNode, error) {
	for node_name, enc_node := range nodes {
		match, _ := regexp.MatchString(node_name, name)
		if match == true {
			return enc_node, nil
		}
	}
	return EncNode{}, errors.New("No node found!")
}
