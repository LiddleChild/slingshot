package parser_test

import (
	"errors"
	"testing"

	"github.com/LiddleChild/slingshot/internal/cli/parser"
	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	t.Run("should error when unknown command", func(t *testing.T) {
		p := parser.NewParser()

		err := p.Parse([]string{})

		assert.Equal(t, errors.New("parser: unknown command"), err, "must return unknown command error")
	})

	t.Run("should error when unknown noun", func(t *testing.T) {
		p := parser.NewParser()

		err := p.Parse([]string{"a"})

		assert.Equal(t, errors.New("parser: unknown noun"), err, "must return unknown noun error")
	})

	t.Run("should error when insufficient arguments", func(t *testing.T) {
		p := parser.NewParser()

		p.Noun("a").Verb("c", func(param *parser.Param) error {
			return nil
		})

		err := p.Parse([]string{"a"})

		assert.Equal(t, errors.New("parser: insufficient arguments"), err, "must return insufficient arguments error")
	})

	t.Run("should error when unknown verb", func(t *testing.T) {
		p := parser.NewParser()

		p.Noun("a").Verb("c", func(param *parser.Param) error {
			return nil
		})

		err := p.Parse([]string{"a", "b"})

		assert.Equal(t, errors.New("parser: unknown verb"), err, "must return unknown verb error")
	})

	t.Run("should success", func(t *testing.T) {
		p := parser.NewParser()

		p.Noun("a").Verb("b", func(param *parser.Param) error {
			return nil
		})

		err := p.Parse([]string{"a", "b"})

		assert.NoError(t, err, "must not return error")
	})
}
