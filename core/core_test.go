package core

import (
	"testing"
)

func TestFilterSlice(t *testing.T) {
	slice := []string{"apple", "banana", "cherry", "date", "elderberry"}

	t.Run("IncludeContains", func(t *testing.T) {
		filtered := FilterSlice(slice, IncludeContains([]string{"a"}))
		expected := []string{"apple", "banana", "date"}
		if !equal(filtered, expected) {
			t.Errorf("Expected %v, got %v", expected, filtered)
		}
	})

	t.Run("ExcludeContains", func(t *testing.T) {
		filtered := FilterSlice(slice, ExcludeContains([]string{"e"}))
		expected := []string{"banana"}
		if !equal(filtered, expected) {
			t.Errorf("Expected %v, got %v", expected, filtered)
		}
	})

	t.Run("IncludeMatches", func(t *testing.T) {
		filtered := FilterSlice(slice, IncludeMatches([]string{"^a", "y$"}))
		expected := []string{"apple", "cherry", "elderberry"}
		if !equal(filtered, expected) {
			t.Errorf("Expected %v, got %v", expected, filtered)
		}
	})

	t.Run("ExcludeMatches", func(t *testing.T) {
		filtered := FilterSlice(slice, ExcludeMatches([]string{"^c", "berry$"}))
		expected := []string{"apple", "banana", "date"}
		if !equal(filtered, expected) {
			t.Errorf("Expected %v, got %v", expected, filtered)
		}
	})

	t.Run("MultipleFilters", func(t *testing.T) {
		filtered := FilterSlice(slice, IncludeContains([]string{"a"}), ExcludeMatches([]string{"^a"}))
		expected := []string{"banana", "date"}
		if !equal(filtered, expected) {
			t.Errorf("Expected %v, got %v", expected, filtered)
		}
	})
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
