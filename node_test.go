package main

import "testing"
import "reflect"
import "regexp"

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

	if node.Environment != "production" {
		t.Error("Expected environment to be Production")
	}

	if node.Parameters["role"] != "foo/bar" {
		t.Error("Expected role to be foo/bar")
	}

	if len(node.Classes) != 2 {
		t.Error("Expected len(node.Classes) to be 2!")
	}

	if !reflect.DeepEqual([]string{"foo::bar", "roles::foo::bar"}, node.Classes) {
		t.Error("Expected node.Classes to have 'foo::bar' and 'roles::foo::bar' only!")
	}
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

	match, _ := regexp.MatchString("classes:\\s+.*\\s*- foo::bar", string_node)
	if match != true {
		t.Error("Expected foo::bar class in stringified node.")
	}

	match, _ = regexp.MatchString("classes:\\s+.*\\s*- roles::foo::bar", string_node)
	if match != true {
		t.Error("Expected roles::foo::bar class in stringified node.")
	}

	match, _ = regexp.MatchString("environment:\\s+production", string_node)
	if match != true {
		t.Error("Unexpected or missing Environment in stringified node.")
	}

	match, _ = regexp.MatchString("parameters:\\s+(\\s+.*)+rack:\\s+r5", string_node)
	if match != true {
		t.Error("Unexpected or missing rack parameter in stringified node.")
	}

	match, _ = regexp.MatchString("parameters:\\s+(\\s+.*)+elevation:\\s+\"5\"", string_node)
	if match != true {
		t.Error("Unexpected or missing elevation parameter found in stringified node.")
	}

	match, _ = regexp.MatchString("parameters:\\s+(\\s+.*)+role:\\s+foo/bar", string_node)
	if match != true {
		t.Error("Unexpected or missing role parameter found in stringified node.")
	}
}
