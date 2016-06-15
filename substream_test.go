package foscam

import (
	"testing"
)

func TestSetSubStreamVideoFormat(t *testing.T) {
	c, err := LoadCredentials("testcredentials.json")
	if err != nil {
		t.Fatal("Unable to load credentials:", err.Error())
	}

	err = SetSubStreamVideoFormat(c, SubStreamVideoFormatMotionJpeg)
	if err != nil {
		t.Fatal(err)
	}
}
