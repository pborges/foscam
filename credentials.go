package foscam
import (
	"os"
	"encoding/json"
)

type Credentials struct {
	Hostname string `json:"hostname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoadCredentials(path string) (c Credentials, err error) {
	fd, err := os.Open(path)
	if err != nil { return }
	err = json.NewDecoder(fd).Decode(&c)
	return
}