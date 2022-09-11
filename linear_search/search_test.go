package linear_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	t.Run("found", func(t *testing.T) {
		result := Search(items, 45)
		assert.Equal(t, 4, result)
	})
	t.Run("found on the first place", func(t *testing.T) {
		result := Search(items, 95)
		assert.Equal(t, 0, result)
	})
	t.Run("not found", func(t *testing.T) {
		result := Search(items, 123)
		assert.Equal(t, -1, result)
	})
}
