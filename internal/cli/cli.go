package cli

import (
	"os"

	"github.com/LiddleChild/slingshot/internal/cli/handler/conn"
	"github.com/LiddleChild/slingshot/internal/cli/parser"
	"github.com/LiddleChild/slingshot/internal/container"
	"github.com/LiddleChild/slingshot/internal/util/logger"
)

func Start(container *container.Container) {
	parser := parser.NewParser()

	connHandler := conn.NewHandler(container.ConnService)

	parser.Noun("conn").
		Verb("list", connHandler.List).
		Verb("create", connHandler.Create)

	err := parser.Parse(os.Args)
	if err != nil {
		logger.Error(err.Error())
	}
}
