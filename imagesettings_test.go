package foscam
import (
	"testing"
)

func TestGetImageSettings(t *testing.T) {
	c, err := LoadCredentials("testcredentials.json")
	if err != nil {
		t.Fatal("Unable to load credentials:", err.Error())
	}

	i, err := GetImageSetting(c)
	if err != nil { t.Fatal(err)}
	if ok, err := i.Success(); !ok {
		t.Fatal(err)
	}
}
