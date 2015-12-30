package main

import (
	"github.com/pborges/foscam"
	"os"
	"image/jpeg"
	"strconv"
)

func main() {
	c := foscam.Credentials{
		Hostname:"nullpntr.com:88",
		Username:"root",
		Password:"Forevertwkd",
	}

	for i := 0; i < 10; i++ {
		img, _ := foscam.SnapPicture2(c)
		f, _ := os.Create("snap" + strconv.Itoa(i) + ".jpg")
		jpeg.Encode(f, img, nil)
		f.Close()
	}
}
