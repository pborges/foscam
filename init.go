package foscam

import (
	"log"
	"os"
)

var DebugLog log.Logger

func init() {
	DebugLog.SetPrefix("[foscam] DEBUG ")
	//DebugLog.SetOutput(ioutil.Discard)
	DebugLog.SetOutput(os.Stdout)
}