package foscam

import (
	"testing"
)

func TestGetWifiConfig(t *testing.T) {
	c, err := LoadCredentials("testcredentials.json")
	if err != nil {
		t.Fatal("Unable to load credentials:", err.Error())
	}

	s, err := GetWifiConfig(c)
	if err != nil {
		t.Fatal(err)
	}
	if ok, err := s.Success(); !ok {
		t.Fatal(err)
	}
}

