package lrucache

import (
	"testing"
)

func TestAdd(t *testing.T) {
	capacity := 100
	c := CreateCache(capacity)


	c.Add("key1", "value1")

	actual := c.cache["key1"].Value.(*entry).value
	expected := "value1"

	if actual != expected {
		t.Errorf("actual %s, expected %s", actual, expected)
	}
}
