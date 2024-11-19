package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func T001_TestServerInitialization(t *testing.T) {
	// Create a new HTTP server using httptest
	server := httptest.NewServer(http.FileServer(http.Dir("./site")))
	defer server.Close()

	// Check if the server is accessible
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to initialize server: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}

func T002_TestFileServing(t *testing.T) {
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

func T003_TestMissingFile(t *testing.T) {
	// Create a new test server
	server := httptest.NewServer(http.FileServer(http.Dir("./site")))
	defer server.Close()

	// Test accessing a non-existent file
	resp, err := http.Get(server.URL + "/nonexistent.html")
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Expected status code 404, got %v", resp.StatusCode)
	}
}

func T004_TestDirectoryBrowsing(t *testing.T) {
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

func T005_TestMimeType(t *testing.T) {
	server := httptest.NewServer(http.FileServer(http.Dir("./site")))
	defer server.Close()

	tests := []struct {
		path        string
		contentType string
	}{
		{"/index.html", "text/html"},
		{"/css/style.css", "text/css"},
		{"/img/logo.png", "image/png"},
	}

	for _, tt := range tests {
		resp, err := http.Get(server.URL + tt.path)
		if err != nil {
			t.Fatalf("Failed to make GET request: %v", err)
		}
		if resp.Header.Get("Content-Type") != tt.contentType {
			t.Errorf("Expected Content-Type %v, got %v for %v", tt.contentType, resp.Header.Get("Content-Type"), tt.path)
		}
	}
}
