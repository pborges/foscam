package foscam

import (
	"image/jpeg"
	"image"
	"io"
	"strconv"
	"encoding/xml"
)

const (
	SnapQualityLow int = 0
	SnapQualityNormal int = 1
	SnapQualityHigh int = 2
)

func SnapPicture(c Credentials) (result io.ReadCloser, err error) {
	params := make(map[string]string)
	params["cmd"] = "snapPicture"

	res, err := request(c, params)
	if err != nil {
		return
	}
	result = res.Body
	return
}

func SnapPicture2(c Credentials) (img image.Image, err error) {
	params := make(map[string]string)
	params["cmd"] = "snapPicture2"

	res, err := request(c, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	img, err = jpeg.Decode(res.Body)
	return
}

func SetSnapConfig(c Credentials, quality int) (err error) {
	params := make(map[string]string)
	params["cmd"] = "setSnapConfig"
	params["snapPicQuality"] = strconv.Itoa(quality)

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
