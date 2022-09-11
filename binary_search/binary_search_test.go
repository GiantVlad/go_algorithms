package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBSearch(t *testing.T) {
	items := []int{1, 3, 9, 34, 45, 86, 99, 251, 320}
	t.Run("found in first iteration", func(t *testing.T) {
		result := BSearch(45, items)
		assert.Equal(t, 4, result)
	})
	t.Run("found", func(t *testing.T) {
		result := BSearch(251, items)
		assert.Equal(t, 7, result)
	})
	t.Run("not found", func(t *testing.T) {
		result := BSearch(95, items)
		assert.Equal(t, -1, result)
	})
}
