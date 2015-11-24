package main

import (
	. "github.com/danzilio/go-yamlenc/Godeps/_workspace/src/github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCollectNodes(t *testing.T) {
	var node_list []string

	node_list = append(node_list, "test/fixtures/node")

	node_list = CollectNodes(node_list)

	Convey("The collection should include nodes.yaml", t, func() {
		So("test/fixtures/node/nodes.yaml", ShouldBeIn, node_list)
	})
}

func TestDir(t *testing.T) {
	node_list := Dir("test/fixtures", "\\.yaml$")

	Convey("The collection should include fixtures/node/nodes.yaml", t, func() {
		So("test/fixtures/node/nodes.yaml", ShouldBeIn, node_list)
	})

	Convey("The collection should include fixtures/node/foo/bar.yaml", t, func() {
		So("test/fixtures/node/foo/bar.yaml", ShouldBeIn, node_list)
	})

	Convey("The collection should not include directories", t, func() {
		So("test/fixtures/node/foo", ShouldNotBeIn, node_list)
	})
}

func TestLookup(t *testing.T) {
	node_list := Dir("test/fixtures/node", "\\.yaml$")

	Convey("It should find dc1-puppet01", t, func() {
		node := Lookup("dc1-puppet01", node_list)
		So(node, ShouldNotBeNil)
		So(node.Role, ShouldEqual, "roles::puppet::master")
		So(node.Environment, ShouldEqual, "production")
	})

	Convey("It should find dc1-puppetdb01.example.com", t, func() {
		node := Lookup("dc1-puppetdb01.example.com", node_list)
		So(node, ShouldNotBeNil)
		So(node.Role, ShouldEqual, "roles::puppet::puppetdb")
		So(node.Environment, ShouldEqual, "stage")
		So("base", ShouldBeIn, node.Classes)
		So("ntp", ShouldBeIn, node.Classes)
	})
}
