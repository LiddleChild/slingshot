package parser_test

import (
	"testing"

	"github.com/LiddleChild/slingshot/internal/cli/parser"
	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	t.Run("HasNext() should return false", func(t *testing.T) {
		p := parser.NewParam([]string{})

		ok := p.HasNext()

		assert.Equal(t, false, ok, "must be false")
	})

	t.Run("Next() should return empty", func(t *testing.T) {
		p := parser.NewParam([]string{})

		result, ok := p.Next()

		assert.Equal(t, false, ok, "must be false")
		assert.Equal(t, "", result, "must be empty string")
	})

	t.Run("Next() should return next", func(t *testing.T) {
		p := parser.NewParam([]string{"a", "b"})

		param, ok := p.Next()

		assert.Equal(t, true, ok, "must be true")
		assert.Equal(t, "a", param, "must be 'a'")
	})
}
