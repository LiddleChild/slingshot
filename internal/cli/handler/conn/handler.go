package conn

import (
	"fmt"

	"github.com/LiddleChild/slingshot/internal/cli/parser"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) List(param *parser.Param) error {
	fmt.Println("list connections")
	return nil
}
