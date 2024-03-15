package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/xagero/go-helper/helper"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	assert.True(t, helper.IsEmpty(""))
}

func TestIsBlank(t *testing.T) {
	assert.True(t, helper.IsBlank(""))
}

func TestIsNotEmpty(t *testing.T) {
	assert.False(t, helper.IsNotEmpty(""))
}

func TestIsNotBlank(t *testing.T) {
	assert.False(t, helper.IsNotBlank(""))
}
