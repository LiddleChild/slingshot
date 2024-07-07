package cli

import (
	"os"

	"github.com/LiddleChild/slingshot/internal/cli/handler/conn"
	"github.com/LiddleChild/slingshot/internal/cli/parser"
	"github.com/LiddleChild/slingshot/internal/util/logger"
)

func Start() {
	parser := parser.NewParser()

	connHandler := conn.NewHandler()

	parser.Noun("conn").Verb("list", connHandler.List)

	err := parser.Parse(os.Args)
	if err != nil {
		logger.Error(err.Error())
	}
}
