package test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xagero/go-helper/helper"
)

var (
	expectedKeys   = []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}
	expectedValues = []string{
		"I am red color!",
		"I am orange color!",
		"I am yellow color!",
		"I am green color!",
		"I am blue color!",
		"I am indigo color!",
		"I am violet color!",
	}
	actualMapTest = map[string]string{
		"red":    "I am red color!",
		"orange": "I am orange color!",
		"yellow": "I am yellow color!",
		"green":  "I am green color!",
		"blue":   "I am blue color!",
		"indigo": "I am indigo color!",
		"violet": "I am violet color!",
	}
)

func TestFetchKeys(t *testing.T) {

	actual := helper.FetchKeys(actualMapTest)

	// Sort data to ensure ordered position
	sort.Strings(expectedKeys)
	sort.Strings(actual)

	assert.Equal(t, expectedKeys, actual)
}

func TestFetchValues(t *testing.T) {
	actual := helper.FetchValues(actualMapTest)

	// Sort data to ensure ordered position
	sort.Strings(expectedValues)
	sort.Strings(actual)

	assert.Equal(t, expectedValues, actual)
}
