package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestNoApiKey(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "")

	_, err := GetAPIKey(h)
	if err == nil {
		t.Fatalf("Function should return error, but it didn't")
	} else if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("Incorrect error returned, expected %q", ErrNoAuthHeaderIncluded)
	}
}

func TestMalformedHeader(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "KeyForApi 65766")

	_, err := GetAPIKey(h)
	if err == nil {
		t.Fatalf("Function should return error, but it didn't")
	} else if errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("Incorrect error returned, expected \"malformed authorization header\"")
	}
}

func TestGoodHeaderNoKey(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "ApiKey")

	_, err := GetAPIKey(h)
	if err == nil {
		t.Fatalf("Function should return error, but it didn't")
	} else if errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("Incorrect error returned, expected \"malformed authorization header\"")
	}
}

func TestProperApiKey(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "ApiKey test123")

	apiKey, err := GetAPIKey(h)
	if err != nil {
		t.Fatalf("Function shouldn't return error, but it did, %v", err)
	}
	if apiKey != "test123" {
		t.Fatalf("Expected \"test123\" API key, got %q instead", apiKey)
	}
}
