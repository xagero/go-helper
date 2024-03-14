package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	assert.Equal(t, Min(1, 2), 1)
	assert.NotEqual(t, Min(1, 2), 2)
	assert.Equal(t, Min(0.1, 0.001), 0.001)
}

func TestMax(t *testing.T) {
	assert.Equal(t, Max(1, 2), 2)
	assert.NotEqual(t, Max(1, 2), 1)
	assert.Equal(t, Max(0.1, 0.001), 0.1)
}
