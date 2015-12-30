package foscam

import "io"

//ToDO: test this
func GetMJStream(c Credentials) (io.Reader, error) {
	params := make(map[string]string)
	params["cmd"] = "GetMJStream"

	res, err := request(c, params)
	if err != nil {
		return nil, err
	}
	return res.Body, err
}

