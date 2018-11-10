package foscam

import (
	"encoding/xml"
	"strconv"
)

const (
	WifiNetTypeInfraNet int = 0
	WifiNetTypeAdHoc int = 1

	WifiEncryptTypeOpen int = 0
	WifiEncryptTypeWep int = 1
	WifiEncryptTypeWpa int = 2
	WifiEncryptTypeWpa2 int = 3
	WifiEncryptTypeWpaWpa2 int = 4

	WifiAuthModeOpen int = 0
	WifiAuthModeSharedKey int = 1
	WifiAuthModeAuto int = 2

	WifiKeyFormatAscii int = 0
	WifiKeyFormatHex int = 1
)

type WifiSettings struct {
	CGIResult
	IsEnable    int `xml:"isEnable"`
	IsUseWifi   int `xml:"isUseWifi"`
	IsConnected int `xml:"isConnected"`
	ConnectedAP string `xml:"connectedAP"`
	SSID        string `xml:"ssid"`
	EncryptType int `xml:"encryptType"`
	Psk         string `xml:"psk"`
	AuthMode    int `xml:"authMode"`
	KeyFormat   int `xml:"keyFormat"`
	DefaultKey  int `xml:"defaultKey"`
	Key1        int `xml:"key1"`
	Key2        int `xml:"key2"`
	Key3        int `xml:"key3"`
	Key4        int `xml:"key4"`
	Key1Len     int `xml:"key1Len"`
	Key2Len     int `xml:"key2Len"`
	Key3Len     int `xml:"key3Len"`
	Key4Len     int `xml:"key4Len"`
}

func GetWifiConfig(c Credentials) (settings *WifiSettings, err error) {
	params := make(map[string]string)
	params["cmd"] = "getWifiConfig"

	res, err := request(c, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	settings = new(WifiSettings)
	xml.NewDecoder(res.Body).Decode(settings)
	return
}

func SetWifiConfig(c Credentials, ssid string, psk string, encryptionType int) (err error) {
	params := make(map[string]string)
	params["cmd"] = "setWifiSetting"
	params["isEnable"] = "1"
	params["isUseWifi"] = "1"
	params["ssid"] = ssid
	params["netType"] = strconv.Itoa(WifiNetTypeInfraNet)
	params["encryptType"] = strconv.Itoa(encryptionType)
	params["psk"] = psk
	params["defaultKey"] = "1"

	res, err := request(c, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	return
}
