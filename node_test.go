package main

import "testing"
import "reflect"

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
