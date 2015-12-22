package foscam

import (
	"encoding/xml"
)

func ChangeUserName(c Credentials, oldUsername string, newUsername string) (err error) {
	params := make(map[string]string)
	params["cmd"] = "changeUserName"
	params["usrName"] = oldUsername
	params["newUsrName"] = newUsername

	res, err := request(c, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	x := new(CGIResult)
	err = xml.NewDecoder(res.Body).Decode(x)
	_, err = x.Success()
	return
}

func ChangePassword(c Credentials, username string, oldPassword string, newPassword string) (err error) {
	params := make(map[string]string)
	params["cmd"] = "changePassword"
	params["usrName"] = username
	params["oldPwd"] = oldPassword
	params["newPwd"] = newPassword

	res, err := request(c, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	x := new(CGIResult)
	err = xml.NewDecoder(res.Body).Decode(x)
	_, err = x.Success()
	return
}