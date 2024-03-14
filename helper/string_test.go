package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	assert.True(t, IsEmpty(""))
}

func TestIsBlank(t *testing.T) {
	assert.True(t, IsBlank(""))
}

func TestIsNotEmpty(t *testing.T) {
	assert.False(t, IsNotEmpty(""))
}

func TestIsNotBlank(t *testing.T) {
	assert.False(t, IsNotBlank(""))
}
