package foscam

func SetOSD(c Credentials, timeStamp bool, devName bool) (err error) {
	params := make(map[string]string)
	params["cmd"] = "setOSDSetting"
	params["isEnableTimeStamp"] = "0"
	if timeStamp {
		params["isEnableTimeStamp"] = "1"
	}
	params["isEnableDevName"] = "0"
	if devName {
                params["isEnableDevName"] = "1"
	}

	res, err := request(c, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	return
}
