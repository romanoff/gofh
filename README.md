Go Flags Handler
================
Go flags handler allows you to handle console commands for your application in similar way as you do it for web applications. You just register handlers and then if pattern of command matches pattern specified for handler, it is called. Go has [flags package](http://golang.org/pkg/flag/), but it is different and doesn't provide routing to different functions based on arguments.

Why would I need it?
--------------------
Let's say you want to implement console utility that will interact with user similarly to git. And you want to start with couple of commands: `project init` and `project deploy`. Here is how your go code would look like:

```go
package main

import (
	"github.com/romanoff/gofh"
	"fmt"
	"os"
)

func main() {
	gofh := gofh.Init()
	gofh.HandleCommand("init", initProject)
	gofh.HandleCommand("deploy", deployProject)
	gofh.SetDefaultHandler(showUsage)
	gofh.Parse(os.Args[1:])
}

func showUsage() {
	fmt.Println("Please, use 'project init' or 'project deploy' command")
}

func initProject(options map[string]string) {
	fmt.Println("Your init project code goes here")
}

func deployProject(options map[string]string) {
	fmt.Println("Your deploy project code goes here")
}
```

Command options
-------------------

You can add a handler with options. Here is an example:

```go
options := []*gofh.Options{
  &Option{Name: "no-css", Boolean: true}
  &Option{Name: "db"}
}

gofh.HandleCommandWithOptions("init", options, initHandler)
```
In this example, no-css is boolean option. So, if you want to supply this option, you just have to add `--no-css` to you console command. It will look like this: `project init --no-css`. After this `initHandler` will get map that will have `no-css` key set to `true`. If `--no-css` option won't be supplied, `no-css` key will be empty.

There is also value option in the above example. to add db option value, you would have to use following console command `project init --db mysql`. In following example `db` key for `initHandler` options would be set to `mysql`.

Summary
-------
I think that this package can be useful to many people. And even though it doesn't cover all possible scenarios for command line arguments, it simplifies creation of git- like command line utilities.
