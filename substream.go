package foscam

import (
	"strconv"
	"encoding/xml"
)

const (
	SubStreamVideoFormatH264 = 0
	SubStreamVideoFormatMotionJpeg = 1
)

func SetSubStreamVideoFormat(c Credentials, videoFormat int) (err error) {
	params := make(map[string]string)
	params["cmd"] = "setSubStreamFormat"
	params["format"] = strconv.Itoa(videoFormat)

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
