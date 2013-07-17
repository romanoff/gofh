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

func TestCommandWithBooleanFlags(t *testing.T) {
	f := Init()
	options := []*Option{
		&Option{Name: "no-views", Boolean: true},
		&Option{Name: "no-components", Boolean: true},
	}
	handlerVisited := false
	f.HandleCommandWithOptions("new :name", options, func(options map[string]string) {
		handlerVisited = true
		if options["name"] != "application" {
			t.Errorf("Expected 'application' argument, but got '%v'", options["name"])
		}
		if options["no-views"] != "true" {
			t.Errorf("Expected 'no-views' argument, but got '%v'", options["no-views"])
		}
	})
	f.Parse([]string{"new", "application", "--no-views"})
	if !handlerVisited {
		t.Error("Did not call callback for command with options handler")
	}
}

func TestCommandWithValueFlags(t *testing.T) {
	f := Init()
	options := []*Option{
		&Option{Name: "db"},
	}
	handlerVisited := false
	f.HandleCommandWithOptions("create :name", options, func(options map[string]string) {
		handlerVisited = true
		if options["db"] != "mysql" {
			t.Errorf("Expected 'db' argument value to be 'mysql', but got '%v'", options["db"])
		}
	})
	f.Parse([]string{"create", "myapp", "--db", "mysql"})
	if !handlerVisited {
		t.Error("Did not call callback for command with value options")
	}
}
