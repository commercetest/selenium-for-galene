package main

import (
    "testing"
    "github.com/tebeka/selenium"
)

func TestBrowserUsingSelenium(t *testing.T) {
    caps := selenium.Capabilities{"browserName": "firefox",
    					"acceptInsecureCerts": true,}
    wd, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
    if err != nil {
        t.Fatalf("Failed to open session: %v", err)
    }
    defer wd.Quit()

    if err := wd.Get("https://localhost:8443"); err != nil {
        t.Fatalf("Failed to load page: %v", err)
    }

    title, err := wd.Title()
    if err != nil {
        t.Fatalf("Failed to get title: %v", err)
    }

    if title != "Galène" {
        t.Errorf("Expected title 'Welcome to My Go Web Server', got %v", title)
    }
}

