package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xagero/go-helper/helper"
)

func TestMin(t *testing.T) {
	assert.Equal(t, helper.Min(1, 2), 1)
	assert.NotEqual(t, helper.Min(1, 2), 2)
	assert.Equal(t, helper.Min(0.1, 0.001), 0.001)
	assert.Equal(t, helper.Min(1, 2, 3, 4, 5, 6, 7, 8, 9), 1)
	assert.Equal(t, helper.Min(1, 2, 3, 4, 5, 4, 5, 2, 1), 1)
	assert.NotEqual(t, helper.Min(1, 2, 3, 4, 5, 4, 5, 2, 1), 5)
}

func TestMax(t *testing.T) {
	assert.Equal(t, helper.Max(1, 2), 2)
	assert.NotEqual(t, helper.Max(1, 2), 1)
	assert.Equal(t, helper.Max(0.1, 0.001), 0.1)
	assert.Equal(t, helper.Max(1, 2, 3, 4, 5, 6, 7, 8, 9), 9)
	assert.Equal(t, helper.Max(1, 2, 3, 4, 5, 4, 3, 2, 1), 5)
	assert.NotEqual(t, helper.Max(1, 2, 3, 4, 5, 4, 3, 2, 1), 1)

}
