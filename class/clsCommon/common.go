package clsCommon

import (
	"math/rand"
	"strconv"
)

func VerifyKey(strKey string) int {
	var i int
	if strKey == "hsulwody638js" {
		i = 1
		return i
	} else {
		i = 0
		return i
	}
}

func VerifyAppVer(strVer string) int {
	var i int
	if strVer == "hsulwody638js" {
		i = 0
		return i
	} else {
		i = 1
		return i
	}
}

func GenerateInt4() string {
	return strconv.FormatInt(int64(rand.Intn(10000)), 10)
}

func GenerateInt6() string {
	return strconv.FormatInt(int64(rand.Intn(1000000)), 10)
}
