package main

import (
	"flag"

	"github.com/LiddleChild/slingshot/internal/cli"
	"github.com/LiddleChild/slingshot/internal/web"
)

var (
	webFlag = flag.Bool("web", false, "open local web ui")
)

func main() {
	flag.Parse()

	if *webFlag {
		web.Start()
	} else {
		cli.Start()
	}
}
