package buble_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	items := []int{-129, 755, 38, 354, 248, 6, -160, 212, -184, 336, -85, 222, 776, 587, 490, -503, -420, -54, -502, -341}
	expected := []int{-503, -502, -420, -341, -184, -160, -129, -85, -54, 6, 38, 212, 222, 248, 336, 354, 490, 587, 755, 776}
	t.Run("sorted", func(t *testing.T) {
		Sort(items)
		assert.Equal(t, expected, items)
	})
}
