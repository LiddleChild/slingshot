package main

import (
	"flag"

	"github.com/LiddleChild/slingshot/internal/cli"
	"github.com/LiddleChild/slingshot/internal/container"
	"github.com/LiddleChild/slingshot/internal/database"
	"github.com/LiddleChild/slingshot/internal/web"
)

var (
	webFlag = flag.Bool("web", false, "open local web ui")
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}

	container := container.NewContainer(db)

	flag.Parse()

	if *webFlag {
		web.Start()
	} else {
		cli.Start(container)
	}
}
