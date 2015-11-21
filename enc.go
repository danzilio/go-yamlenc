package main

import "os"
import "github.com/danzilio/go-yamlenc/Godeps/_workspace/src/github.com/codegangsta/cli"

func ParseOpts(arguments []string) {
	app := cli.NewApp()
	app.Name = "yamlenc"
	app.Usage = "A very simple external node classifier (ENC) for Puppet"
	app.Version = Version

	var fail bool
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
		for _, element := range cli_nodes {
			config.NodeList = append(config.NodeList, element)
		}
	}

	app.Run(arguments)
}

func main() {
	ParseOpts(os.Args)
}
