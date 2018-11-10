package foscam

import (
	"encoding/xml"
	"strconv"
	"errors"
)

type AlarmRecordConfig struct {
	IsEnablePreRecord  int `xml:"isEnablePreRecord"`
	PreRecordSecs      int `xml:"preRecordSecs"`
	AlarmRecordSeconds int `xml:"alarmRecordSeconds"`
}

func GetAlarmRecordConfig(c Credentials) (settings *AlarmRecordConfig, err error) {
	params := make(map[string]string)
	params["cmd"] = "getAlarmRecordConfig"

	res, err := request(c, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	settings = new(AlarmRecordConfig)
	xml.NewDecoder(res.Body).Decode(settings)
	return
}

func SetAlarmRecordConfig(c Credentials, settings AlarmRecordConfig) (err error) {
	// validate AlarmRecordConfig
	if !(settings.IsEnablePreRecord >= 0 && settings.IsEnablePreRecord <= 1) {
		err = errors.New("IsEnablePreRecord must be 0 (disabled) or 1 (enabled)")
		return
	}
	if !(settings.PreRecordSecs > 0 && settings.PreRecordSecs <= 5) {
		err = errors.New("PreRecordSecs must be between 1 and 5")
		return
	}
	switch settings.AlarmRecordSeconds {
	case 10, 20, 30, 40, 50, 60, 120, 180, 240, 300:
	default:
		err = errors.New("AlarmRecordSeconds must be 10,20,30,40,50,60,120,180,240,300")
		return
	}
	params := make(map[string]string)
	params["cmd"] = "setAlarmRecordConfig"
	params["isEnablePreRecord"] = strconv.Itoa(settings.IsEnablePreRecord)
	params["preRecordSecs"] = strconv.Itoa(settings.PreRecordSecs)
	params["alarmRecordSecs"] = strconv.Itoa(settings.AlarmRecordSeconds)

	res, err := request(c, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	result := new(CGIResult)
	xml.NewDecoder(res.Body).Decode(result)
	_, err = result.Success()
	return
}