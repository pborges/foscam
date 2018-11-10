package foscam

import "testing"

func TestChangeUserName(t *testing.T) {
	c, err := LoadCredentials("testcredentials.json")
	if err != nil {
		t.Fatal("Unable to load credentials:", err.Error())
	}

	err = ChangeUserName(c, "test", "test2")
	if err != nil {
		t.Fatal("Unable to change username " + err.Error())
	}

	err = ChangeUserName(c, "test2", "test")
	if err != nil {
		t.Fatal("Unable to change username back " + err.Error())
	}
}

func TestChangePassword(t *testing.T) {
	c, err := LoadCredentials("testcredentials.json")
	if err != nil {
		t.Fatal("Unable to load credentials:", err.Error())
	}

	err = ChangePassword(c, "test", "test", "test2")
	if err != nil {
		t.Fatal("Unable to change password " + err.Error())
	}

	err = ChangePassword(c, "test", "test2", "test")
	if err != nil {
		t.Fatal("Unable to change password back " + err.Error())
	}
}
