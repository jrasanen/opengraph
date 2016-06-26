package main

import "testing"

func expectToBeEqual(t *testing.T, expected string, result string) {
	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func expectToBeEqualOrGreater(t *testing.T, expected int, result int) {
	if expected > result {
		t.Errorf("Expected at least %v, got %v", expected, result)
	}
}

func TestGetTags(t *testing.T) {
	tags := GetTags("https://twitter.com/ID_AA_Carmack/status/747089811853348864")

	expectToBeEqualOrGreater(t, 1, len(tags.GetTagsByName("og:url")))
	expectToBeEqualOrGreater(t, 1, len(tags.GetTagsByName("og:image")))
}
