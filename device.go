package foscam

import (
	"encoding/xml"
)

type DeviceInfo struct {
	CGIResult
	ProductName  string `xml:"productName" json:"productName"`
	SerialNumber string `xml:"serialNo" json:"serialNo"`
	DeviceName   string `xml:"devName" json:"devName"`
	MacAddress   string `xml:"mac" json:"mac"`
	Year         string `xml:"year" json:"year"`
	Month        string `xml:"mon" json:"mon"`
}

func SetDevName(c Credentials, name string) (err error) {
	params := make(map[string]string)
	params["cmd"] = "setDevName"
	params["devName"] = name

	res, err := request(c, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	x := new(CGIResult)
	err = xml.NewDecoder(res.Body).Decode(x)
	if err != nil {
		return
	}
	_, err = x.Success()
	return
}

func GetDevName(c Credentials) (name string, err error) {
	params := make(map[string]string)
	params["cmd"] = "getDevName"

	res, err := request(c, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	x := new(DeviceInfo)
	err = xml.NewDecoder(res.Body).Decode(x)
	if err != nil {
		return
	}
	_, err = x.Success()
	if err != nil {
		return
	}
	name = x.DeviceName
	return
}

func GetDevInfo(c Credentials) (info *DeviceInfo, err error) {
	params := make(map[string]string)
	params["cmd"] = "getDevInfo"

	res, err := request(c, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	info = new(DeviceInfo)
	xml.NewDecoder(res.Body).Decode(info)
	_, err = info.Success()
	return
}

//ToDo: SetSystemTime
//ToDo: GetSystemTime
//ToDo: getOSDSetting
//ToDo: SetOSDSetting