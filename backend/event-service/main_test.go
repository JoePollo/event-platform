package main

import (
	"net/http"
	"testing"
)

func TestConnectToDb(t *testing.T) {
	url := "http://localhost:8080/connectToDb"
	_, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("Failed to connect to database due to error:\n%v", err)
	}
}
