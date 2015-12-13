package foscam
import (
	"testing"
	"io/ioutil"
	"regexp"
)


func TestSnapPicture(t *testing.T) {
	c, err := LoadCredentials("testcredentials.json")
	if err != nil {t.Fatal("Unable to load credentials:", err.Error())}

	r, err := SnapPicture(c)
	if err != nil {
		t.Fatal("Unable to SnapPicture", err.Error())
	}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal("Unable to read body", err.Error())
	}
	ok, err := regexp.Match(`<html><body><img src="\.\./snapPic/Snap_\d+-\d+\.jpg"/></body></html>`, data)
	if err != nil {
		t.Fatal("Unable to perform regexpy", err.Error())
	}
	if !ok {
		t.Fatal("No match found", string(data))
	}
}

func TestSnapPicture2(t *testing.T) {
	c, err := LoadCredentials("testcredentials.json")
	if err != nil {t.Fatal("Unable to load credentials:", err.Error())}

	r, err := SnapPicture2(c)
	if err != nil {
		t.Fatal(err.Error())
	}
	if r.Bounds().Dx() != 1280 {
		t.Fatal("Invalid Width, expected 1280 got", r.Bounds().Dx())
	}
	if r.Bounds().Dy() != 720 {
		t.Fatal("Invalid Height, expected 720 got", r.Bounds().Dy())
	}
}