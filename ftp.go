package foscam

import (
	"encoding/xml"
	"strconv"
)

const (
	FtpModePasv int = 0
	FtpModePort int = 1
)

type FtpTestResult struct {
	CGIResult
	TestResult int `xml:"testResult"`
}

type FtpSettings struct {
	CGIResult
	Result   string `xml:"result"`
	Address  string `xml:"ftpAddr"`
	Port     int `xml:"ftpPort"`
	Mode     int `xml:"mode"`
	UserName string `xml:"userName"`
	Password string `xml:"password"`
}

func GetFtpConfig(c Credentials) (settings *FtpSettings, err error) {
	params := make(map[string]string)
	params["cmd"] = "getFtpConfig"

	res, err := request(c, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	settings = new(FtpSettings)
	xml.NewDecoder(res.Body).Decode(settings)
	_, err = settings.Success()
	return
}

func SetFtpConfig(c Credentials, settings FtpSettings) (err error) {
	params := make(map[string]string)
	params["cmd"] = "setFtpConfig"
	params["ftpAddr"] = settings.Address
	params["ftpPort"] = strconv.Itoa(settings.Port)
	params["mode"] = strconv.Itoa(settings.Mode)
	params["userName"] = settings.UserName
	params["password"] = settings.Password

	res, err := request(c, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	result := new(DeviceInfo)
	xml.NewDecoder(res.Body).Decode(result)
	_, err = result.Success()
	return
}

func TestFtpServer(c Credentials, settings FtpSettings) (ok bool, err error) {
	params := make(map[string]string)
	params["cmd"] = "testFtpServer"
	params["ftpAddr"] = settings.Address
	params["ftpPort"] = strconv.Itoa(settings.Port)
	params["mode"] = strconv.Itoa(settings.Mode)
	params["userName"] = settings.UserName
	params["password"] = settings.Password

	res, err := request(c, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	result := new(FtpTestResult)
	xml.NewDecoder(res.Body).Decode(result)
	ok, err = result.Success()
	if err != nil {
		return
	}
	ok = result.TestResult == 0
	return
}