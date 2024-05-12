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


func TestGet(t *testing.T){
	t.Run("get an item from the cache when the key is supplied", func(t *testing.T) {
		capacity := 2
		c := CreateCache(capacity)

		c.Add("key1", "value1")
		c.Add("key2", "value2")
		c.Add("key3", "value3")

		actual := c.Get("key2")
		expected := "value2"

		if actual != expected {
			t.Errorf("actual %s, expected %s", actual, expected)
		}
	})
}