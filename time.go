package foscam

import "encoding/xml"

func SetSystemTimeNTP(c Credentials, dateFormat string, ntpServer string, timeZoneOffset int, use24hour bool, dst bool) (err error) {
	params := make(map[string]string)
	params["cmd"] = "setSystemTime"
	params["dateFormat"] = dateFormat
	params["timeSource"] = strconv.Itoa(1)
	params["ntpServer"] = ntpServer
	params["timeZone"] = timeZoneOffset * -1 * 3600
	if dst {
		params["isDst"] = 1
	} else {
		params["isDst"] = 0
	}

	if use24hour {
		params["timeFormat"] = 1
	} else {
		params["timeFormat"] = 0
	}

	res, err := request(c, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	x := new(CGIResult)
	err = xml.NewDecoder(res.Body).Decode(x)
	if err != nil {
		return
	}
	_, err = x.Success()
	return
}
//ToDo: GetSystemTime


const TimeDateFormatYYYYMMDD = "YYYY-MM-DD"
const TimeDateFormatDDMMYYYY = "DD/MM/YYYY"
const TimeDateFormatMMDDYYYY = "MM/DD/YYYY"