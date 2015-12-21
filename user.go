package foscam

func ChangeUserNameAndPwdTogether(c Credentials, oldUsername string, oldPassword string, newUsername string, newPassword string) (err error) {
	params := make(map[string]string)
	params["cmd"] = "changeUserNameAndPwdTogether"
	params["usrName"] = oldUsername
	params["newUsrName"] = newUsername
	params["oldPwd"] = oldPassword
	params["newPwd"] = newPassword

	res, err := request(c, params)
	if err != nil {
		return
	}
	defer res.Body.Close()
	return
}