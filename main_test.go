package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerInitialization(t *testing.T) {
	// Create a new HTTP server using httptest
	fs := http.FileServer(http.Dir("./site"))
	testserver := httptest.NewServer(fs)
	defer testserver.Close()
	url := testserver.URL + "/"

	// Check if the server is accessible
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("Failed to initialize server: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", err)
	}
}

func TestFileServing(t *testing.T) {
	// Create a new test server
	server := httptest.NewServer(http.FileServer(http.Dir("./site")))
	defer server.Close()

	// Test serving an existing file
	resp, err := http.Get(server.URL + "/index.html")
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}

func TestDirectoryBrowsing(t *testing.T) {
	// Create a new test server
	server := httptest.NewServer(http.FileServer(http.Dir("./site")))
	defer server.Close()

	// Test directory browsing
	resp, err := http.Get(server.URL + "/css/")
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("Expected directory browsing to be disabled, but it succeeded")
	}
}
