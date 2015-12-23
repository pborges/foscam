package foscam

import "testing"

func TestGetAlarmRecordConfig(t *testing.T) {
	c, err := LoadCredentials("testcredentials.json")
	if err != nil {
		t.Fatal("Unable to load credentials:", err.Error())
	}

	_, err = GetAlarmRecordConfig(c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetAlarmRecordConfig(t *testing.T) {
	c, err := LoadCredentials("testcredentials.json")
	if err != nil {
		t.Fatal("Unable to load credentials:", err.Error())
	}

	err = SetAlarmRecordConfig(c, AlarmRecordConfig{
		IsEnablePreRecord:0,
		PreRecordSecs:1,
		AlarmRecordSeconds:20,
	})
	if err != nil {
		t.Fatal(err)
	}
}