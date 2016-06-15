package foscam

func DefaultCredentials(addr string) (Credentials) {
	return Credentials{
		Hostname:addr,
		Username:"admin",
		Password:"",
	}
}
