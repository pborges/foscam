package main

import (
	"github.com/pborges/foscam"
	"fmt"
)

func main() {
	c := foscam.Credentials{
		Hostname:"192.168.2.53:88",
		Username:"root",
		Password:"Forevertwkd",
	}

	f := foscam.FtpSettings{
		Address:"ftp://192.168.2.102",
		Port:2121,
		Mode:foscam.FtpModePasv,
		UserName:"test",
		Password:"pass",
	}

	err := foscam.SetFtpConfig(c, f)
	if err != nil {
		fmt.Println(err)
	}

	ok, err := foscam.TestFtpServer(c, f)
	fmt.Println(ok)
}
