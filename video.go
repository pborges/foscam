package foscam

import (
	"strconv"
	"encoding/xml"
	"errors"
)

const VideoResolution720P = 0
const VideoResolutionVGA640x480 = 1
const VideoResolutionVGA640x360 = 3
const VideoResolutionQVGA320x240 = 2
const VideoResolutionQVGA320x180 = 4

const VideoBitRate4M = 4194304
const VideoBitRate2M = 2097152
const VideoBitRate1M = 1048576
const VideoBitRate512K = 524288
const VideoBitRate200K = 204800
const VideoBitRate100K = 102400
const VideoBitRate50K = 51200
const VideoBitRate20K = 20480

func SetVideoStream(c Credentials, streamType int, resolution int, bitRate int, frameRate int, keyFrameInterval int, variableBitRate bool) (err error) {
	if !(streamType >= 0 && streamType <= 4) {
		err = errors.New("Stream Type must be between 0 and 4")
		return
	}
	if !(frameRate >= 1 && frameRate <= 30) {
		err = errors.New("Frame Rate must be between 1 and 30")
		return
	}
	if !(keyFrameInterval >= 10 && keyFrameInterval <= 30) {
		err = errors.New("Key Frame Interval must be between 10 and 100")
		return
	}

	params := make(map[string]string)
	params["cmd"] = "setVideoStreamParam"
	params["streamType"] = strconv.Itoa(streamType)
	params["resolution"] = strconv.Itoa(resolution)
	params["bitRate"] = strconv.Itoa(bitRate)
	params["frameRate"] = strconv.Itoa(frameRate)
	params["GOP" ] = strconv.Itoa(keyFrameInterval)
	if variableBitRate {
		params["isVBR"] = "1"
	} else {
		params["isVBR"] = "0"
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

