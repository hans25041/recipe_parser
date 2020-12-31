package recipe_parser

import (
	"testing"
)

func assertStringsEqual(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected: %q Got: %q", expected, actual)
	}
}

func assertErrorsEqual(t testing.TB, actual error, expected HostNameNotMinimalistBakerError) {
	t.Helper()
	if actual != expected {
		t.Errorf("Got %q but expected %q", actual, expected)
	}
}

func expectPanic(t testing.TB) {
	t.Helper()
	if r := recover(); r == nil {
		t.Errorf("The code did not panic.")
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Unexpected error returned: %v", err)
	}
}

func TestMinimalistBakerUrlParser(t *testing.T) {
	t.Run("Test panic if the URL isn't a valid URL.", func(t *testing.T) {
		defer expectPanic(t)
		_, err := ParseUrl("not_a_url")
		assertNoError(t, err)
	})

	t.Run("Test parse a URL with the wrong hostname.", func(t *testing.T) {
		_, err := ParseUrl("https://www.epicurious.com/some/recipe/")
		expectedErr := HostNameNotMinimalistBakerError("epicurious.com")
		assertErrorsEqual(t, err, expectedErr)
	})

	t.Run("Test parse a URL with the correct hostname.", func(t *testing.T) {
		url, _ := ParseUrl("https://minimalistbaker.com/easy-vegan-fried-rice/")

		actualHostname := url.Hostname()
		assertStringsEqual(t, actualHostname, "minimalistbaker.com")

		actualPath := url.Path
		assertStringsEqual(t, actualPath, "/easy-vegan-fried-rice/")
	})
}
