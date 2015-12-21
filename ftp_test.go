package foscam

import (
	"testing"
)

func TestGetFtpConfig(t *testing.T) {
	c, err := LoadCredentials("testcredentials.json")
	if err != nil {
		t.Fatal("Unable to load credentials:", err.Error())
	}

	i, err := GetFtpConfig(c)
	if err != nil {
		t.Fatal(err)
	}
	if ok, err := i.Success(); !ok {
		t.Fatal(err)
	}
}

func TestSetFtpConfig(t *testing.T) {
	c, err := LoadCredentials("testcredentials.json")
	if err != nil {
		t.Fatal("Unable to load credentials:", err.Error())
	}

	err = SetFtpConfig(c, FtpSettings{
		Address:"127.0.0.1",
		Port:21,
		Mode:FtpModePasv,
		UserName:"test_user",
		Password:"test_pass",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestTestFtpServer(t *testing.T) {
	c, err := LoadCredentials("testcredentials.json")
	if err != nil {
		t.Fatal("Unable to load credentials:", err.Error())
	}

	_, err = TestFtpServer(c, FtpSettings{
		Address:"0.0.0.0",
		Port:21,
		Mode:FtpModePasv,
		UserName:"test_user",
		Password:"test_pass",
	})
	if err != nil {
		t.Fatal(err)
	}

}
