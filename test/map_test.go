package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/xagero/go-helper/helper"
	"sort"
	"testing"
)

func TestFetchKeys(t *testing.T) {

	colors := []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}

	rainbow := make(map[string]string)

	rainbow["red"] = "I am red color!"
	rainbow["orange"] = "I am orange color!"
	rainbow["yellow"] = "I am yellow color!"
	rainbow["green"] = "I am green color!"
	rainbow["blue"] = "I am blue color!"
	rainbow["indigo"] = "I am indigo color!"
	rainbow["violet"] = "I am violet color!"

	keys := helper.FetchKeys(rainbow)

	// Sort data to ensure ordered position
	sort.Strings(colors)
	sort.Strings(keys)

	assert.Equal(t, colors, keys)
}
