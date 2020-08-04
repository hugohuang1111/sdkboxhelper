package env

import (
	"path/filepath"
	"runtime"
	"errors"

	"sdkboxhelper/utils"
)

var sdkboxHome string

// UserHome user home
func UserHome() string {
	userHome, err := utils.UserHome()
	if nil != err {
		panic(err)
	}
	return userHome
}

// SDKBoxHome sdkbox home
func SDKBoxHome() string {
	if "" != sdkboxHome {
		return sdkboxHome
	}
	userHome, err := utils.UserHome()
	if nil != err {
		panic(err)
	}
	sdkboxHome = filepath.Join(userHome, ".sdkbox")
	return sdkboxHome
}

// GetEnvVar get env variable
func GetEnvVar(key string) (string, bool) {

	return "", false
}

// AddEnvVar add env variable
func AddEnvVar(key string, value string) error {
	switch runtime.GOOS {
	case "windows":
		return addWinEnvVar(key, value)
	case "linux", "darwin":
		return addUnixEnvVar(key, value)
	default:
		return errors.New("unsupport platform:" + runtime.GOOS)
	}
}
