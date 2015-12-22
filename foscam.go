package foscam

import (
	"net/http"
	"strings"
	"encoding/xml"
	"errors"
	"strconv"
	"net/url"
)

const basePath string = "/cgi-bin/CGIProxy.fcgi"

func request(c Credentials, params map[string]string) (*http.Response, error) {
	paramsArr := make([]string, 0)
	params["usr"] = c.Username
	params["pwd"] = c.Password

	for key, value := range params {
		paramsArr = append(paramsArr, key + "=" + url.QueryEscape(value))
	}
	url := "http://" + c.Hostname + basePath + "?" + strings.Join(paramsArr, "&")
	DebugLog.Printf("GET %s\n", url)
	res, err := http.Get(url)
	if err == nil && res.StatusCode != 200 {
		return res, errors.New("Bad Resposne Code, " + strconv.Itoa(res.StatusCode))
	}
	return res, err
}

type CGIResult struct {
	XMLName xml.Name `xml:"CGI_Result"`
	Result  int `xml:"result"`
}

func (c CGIResult) Success() (ok bool, err error) {
	switch c.Result {
	case 0:
		ok = true
	case -1:
		err = errors.New("CGI request string format error")
	case -2:
		err = errors.New("Username or password error")
	case -3:
		err = errors.New("Access deny")
	case -4:
		err = errors.New("CGI execute fail")
	case -5:
		err = errors.New("Timeout")
	case -6:
		err = errors.New("Reserve")
	case -7:
		err = errors.New("Unknown error")
	case -8:
		err = errors.New("Reserve")
	}
	return
}