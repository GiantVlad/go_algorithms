package interpolation_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	items := []int{2, 5, 6, 8, 12, 18, 45, 111, 251, 320}
	t.Run("found", func(t *testing.T) {
		result := InterpolSearch(45, items)
		assert.Equal(t, 6, result)
	})
	t.Run("found on the first place", func(t *testing.T) {
		result := InterpolSearch(2, items)
		assert.Equal(t, 0, result)
	})
	t.Run("found on the last place", func(t *testing.T) {
		result := InterpolSearch(320, items)
		assert.Equal(t, 9, result)
	})
	t.Run("not found", func(t *testing.T) {
		result := InterpolSearch(123, items)
		assert.Equal(t, -1, result)
	})
	t.Run("not found less", func(t *testing.T) {
		result := InterpolSearch(1, items)
		assert.Equal(t, -1, result)
	})
	t.Run("not found more", func(t *testing.T) {
		result := InterpolSearch(321, items)
		assert.Equal(t, -1, result)
	})
}
