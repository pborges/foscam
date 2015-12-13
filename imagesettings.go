package foscam
import (
	"encoding/xml"
)

type ImageSetting struct {
	CGIResult
	Brightness int  `xml:"brightness"`
	Contrast   int  `xml:"contrast"`
	Hue        int  `xml:"hue"`
	Saturation int  `xml:"saturation"`
	Sharpness  int  `xml:"sharpness"`
}

func GetImageSetting(c Credentials) (settings *ImageSetting, err error) {
	params := make(map[string]string)
	params["cmd"] = "getImageSetting"

	res, err := request(c, params)
	if err != nil { return }
	defer res.Body.Close()
	settings = new(ImageSetting)
	xml.NewDecoder(res.Body).Decode(settings)
	return
}
