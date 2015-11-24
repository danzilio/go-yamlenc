package main

import (
	. "github.com/danzilio/go-yamlenc/Godeps/_workspace/src/github.com/smartystreets/goconvey/convey"
	"regexp"
	"testing"
)

func TestPuppetNode(t *testing.T) {
	params := make(map[string]string)
	params["rack"] = "r5"
	params["elevation"] = "5"
	subject := EncNode{
		Role:        "roles::foo::bar",
		Environment: "production",
		Classes:     []string{"foo::bar"},
		Parameters:  params,
	}

	node := subject.ToPuppetNode()

	Convey("The node's environment should be production", t, func() {
		So(node.Environment, ShouldEqual, "production")
	})

	Convey("The node's role should be foo/bar", t, func() {
		So(node.Parameters["role"], ShouldEqual, "foo/bar")
	})

	Convey("The node should have two classes", t, func() {
		So(len(node.Classes), ShouldEqual, 2)
	})

	Convey("The node should have foo::bar in its class array", t, func() {
		So("foo::bar", ShouldBeIn, node.Classes)
	})

	Convey("The node should have roles::foo::bar in its class array", t, func() {
		So("roles::foo::bar", ShouldBeIn, node.Classes)
	})
}

func TestString(t *testing.T) {
	params := make(map[string]string)
	params["rack"] = "r5"
	params["elevation"] = "5"
	subject := EncNode{
		Role:        "roles::foo::bar",
		Environment: "production",
		Classes:     []string{"foo::bar"},
		Parameters:  params,
	}

	node := subject.ToPuppetNode()
	string_node := node.String()

	Convey("The node should have foo::bar in its class array", t, func() {
		match, _ := regexp.MatchString("classes:\\s+.*\\s*- foo::bar", string_node)
		So(match, ShouldBeTrue)
	})

	Convey("The node should have roles::foo::bar in its class array", t, func() {
		match, _ := regexp.MatchString("classes:\\s+.*\\s*- roles::foo::bar", string_node)
		So(match, ShouldBeTrue)
	})

	Convey("The node's environment should be production", t, func() {
		match, _ := regexp.MatchString("environment:\\s+production", string_node)
		So(match, ShouldBeTrue)
	})

	Convey("The node's rack paramter should be set to r5", t, func() {
		match, _ := regexp.MatchString("parameters:\\s+(\\s+.*)+rack:\\s+r5", string_node)
		So(match, ShouldBeTrue)
	})

	Convey("The node's elevation parameter should be set to 5", t, func() {
		match, _ := regexp.MatchString("parameters:\\s+(\\s+.*)+elevation:\\s+\"5\"", string_node)
		So(match, ShouldBeTrue)
	})

	Convey("The node's role parameter should be set to foo/bar", t, func() {
		match, _ := regexp.MatchString("parameters:\\s+(\\s+.*)+role:\\s+foo/bar", string_node)
		So(match, ShouldBeTrue)
	})
}
