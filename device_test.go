package foscam

import "testing"

func TestSetDevName(t *testing.T) {
	c, err := LoadCredentials("testcredentials.json")
	if err != nil {
		t.Fatal("Unable to load credentials:", err.Error())
	}

	err = SetDevName(c, "test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetDevName(t *testing.T) {
	c, err := LoadCredentials("testcredentials.json")
	if err != nil {
		t.Fatal("Unable to load credentials:", err.Error())
	}

	name, err := GetDevName(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Got Device Name " + name)
}

func TestGetDevInfo(t *testing.T) {
	c, err := LoadCredentials("testcredentials.json")
	if err != nil {
		t.Fatal("Unable to load credentials:", err.Error())
	}

	info, err := GetDevInfo(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Got Device Name %+v\n", info)
}
//ToDo: TestSetSystemTime
//ToDo: TestGetSystemTime
//ToDo: TestgetOSDSetting
//ToDo: TestSetOSDSetting