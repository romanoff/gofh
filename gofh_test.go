package gofh

import (
	"testing"
)

func TestFlags(t *testing.T) {
	f := Init()
	initCalled := false
	f.HandleCommand("init", func(options map[string]string) {
		initCalled = true
	})
	f.Parse([]string{"init"})
	if !initCalled {
		t.Error("Expected init callback to be called")
	}
}

func TestFlagsWithArguments(t *testing.T) {
	f := Init()
	f.HandleCommand("new :name", func(options map[string]string) {
		if options["name"] != "application" {
			t.Errorf("Expected 'application' argument, but got '%v'", options["name"])
		}
	})
	f.Parse([]string{"new", "application"})
}

func TestDefaultHandler(t *testing.T) {
	f := Init()
	f.HandleCommand("new :name", func(options map[string]string) {
		t.Error("Not expected to handle new command")
	})
	visitedDefaultHandler := false
	f.SetDefaultHandler(func() {
		visitedDefaultHandler = true
	})
	f.Parse([]string{"help"})
	if !visitedDefaultHandler {
		t.Error("Did not call callback for default handler")
	}
}