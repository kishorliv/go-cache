package lrucache

import (
	"testing"
)

func TestAdd(t *testing.T) {

	t.Run("add item in the cache", func(t *testing.T) {
		capacity := 100
		c := CreateCache(capacity)

		c.Add("key1", "value1")

		actual := c.cache["key1"].Value.(*entry).value
		expected := "value1"

		if actual != expected {
			t.Errorf("actual %s, expected %s", actual, expected)
		}
	})

	t.Run("should evict the item from the tail and add the new item at the head when the cache is full", func(t *testing.T) {
		capacity := 2
		c := CreateCache(capacity)

		c.Add("key1", "value1")
		c.Add("key2", "value2")

		actual := c.cache["key1"].Value.(*entry).value
		expected := "value1"

		if actual != expected {
			t.Errorf("actual %s, expected %s", actual, expected)
		}

		actual = c.cache["key2"].Value.(*entry).value
		expected = "value2"

		if actual != expected {
			t.Errorf("actual %s, expected %s", actual, expected)
		}

		c.Add("key3", "value3")

		if c.list.Len() != capacity {
			t.Errorf("incorrect cache list length")
		}

		if c.list.Front().Value.(*entry).value != "value3" {
			t.Errorf("incorrect item or item position in the cache")
		}

		if c.list.Front().Next().Value.(*entry).value != "value2" {
			t.Errorf("incorrect item or item position in the cache")
		}
	})
}

func TestGet(t *testing.T) {
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
