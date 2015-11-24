package main

import (
	. "github.com/danzilio/go-yamlenc/Godeps/_workspace/src/github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLoad(t *testing.T) {
	Convey("It should load the config with nodes set to a string", t, func() {
		config := Config{}
		config.Load("test/fixtures/conf/string_node.yaml")
		So(config.NodeList, ShouldNotBeEmpty)
		So(config.NodeList, ShouldContain, "/tmp/aruba/nodes")
	})

	Convey("It should load the config with nodes set to a slice", t, func() {
		config := Config{}
		config.Load("test/fixtures/conf/slice_node.yaml")
		So(config.NodeList, ShouldNotBeEmpty)
		So(config.NodeList, ShouldContain, "/tmp/aruba/nodes")
	})
}
