package recipe_parser

import "testing"

func assertStringsEqual(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected: %q Got: %q", expected, actual)
	}
}

func TestMinimalistBakerParser(t *testing.T) {
	t.Run("Test Hello", func(t *testing.T) {
		actual := Hello()
		expected := "Hello"
		assertStringsEqual(t, actual, expected)
	})
}
