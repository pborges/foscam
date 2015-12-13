package foscam
import (
	"image/jpeg"
	"image"
	"io"
)

func SnapPicture(c Credentials) (result io.ReadCloser, err error) {
	params := make(map[string]string)
	params["cmd"] = "snapPicture"

	res, err := request(c, params)
	if err != nil { return }
	result = res.Body
	return
}


func SnapPicture2(c Credentials) (img image.Image, err error) {
	params := make(map[string]string)
	params["cmd"] = "snapPicture2"

	res, err := request(c, params)
	if err != nil { return }
	defer res.Body.Close()
	img, err = jpeg.Decode(res.Body)
	return
}
