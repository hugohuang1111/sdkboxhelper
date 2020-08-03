package env

var sdkboxHome string

func UserHome() string {
	userHome, err := utils.Home()
	if nil != err {
		panic(err)
	}
	return userHome
}

func SDKBoxHome() string {
	if "" != sdkboxHome {
		return sdkboxHome
	}
	userHome, err := utils.Home()
	if nil != err {
		panic(err)
	}
	sdkboxHome = filepath.Join(userHome, ".sdkbox")
	return sdkboxHome
}

