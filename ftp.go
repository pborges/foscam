package foscam

import (
	"encoding/xml"
	"strconv"
	"io/ioutil"
)

const (
	FtpModePasv int = 0
	FtpModePort int = 1
)

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
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	ok = (string(result) == "0")
	return
}