package cache

import (
	"testing"
	"time"
)

func TestPut(t *testing.T) {
	cache := NewCache()
	cache.Put("1", "One")
	cache.Put("2", "Two")
	cache.Put("3", "Three")

	want := 3
	got := len(cache.Keys())

	if got != want {
		t.Errorf("TestPut got: %v, want: %v", got, want)
	}
}

func TestGetElement(t *testing.T) {
	cache := NewCache()
	cache.Put("1", "One")
	cache.Put("2", "Two")
	cache.Put("3", "Three")

	want := "Two"
	got, _ := cache.Get("2")

	if got != want {
		t.Errorf("TestGetElement got: %v, want: %v", got, want)
	}
}

func TestPutTillWithDeadlineNoItems(t *testing.T) {
	cache := NewCache()
	cache.PutTill("NoShow", "no show", time.Date(2020, time.June, 21, 20, 11, 0, 0, time.Local))

	want := 0
	got := len(cache.Keys())

	if got != want {
		t.Errorf("TestPutTillWithDeadlineNoItems got: %v, want: %v", got, want)
	}
}

func TestPutTillWithDeadlineOneItem(t *testing.T) {
	cache := NewCache()
	cache.PutTill("NoShow", "no show", time.Date(2020, time.June, 21, 20, 11, 0, 0, time.Local))
	cache.PutTill("ShowOne", "Show one", time.Date(2023, time.June, 21, 15, 57, 0, 0, time.Local))

	want := 1
	got := len(cache.Keys())

	if got != want {
		t.Errorf("TestPutTillWithDeadlineOneItem got: %v, want: %v", got, want)
	}
}
