package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/xagero/go-helper/helper"
	"testing"
)

func TestMin(t *testing.T) {
	assert.Equal(t, helper.Min(1, 2), 1)
	assert.NotEqual(t, helper.Min(1, 2), 2)
	assert.Equal(t, helper.Min(0.1, 0.001), 0.001)
}

func TestMax(t *testing.T) {
	assert.Equal(t, helper.Max(1, 2), 2)
	assert.NotEqual(t, helper.Max(1, 2), 1)
	assert.Equal(t, helper.Max(0.1, 0.001), 0.1)
}
