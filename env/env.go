package env

import (
	"path/filepath"

	"sdkbox.com/helper/utils"
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
