package installer

import (
	"fmt"
	"path/filepath"

	"sdkbox.com/helper/env"
	"sdkbox.com/helper/url"
	"sdkbox.com/helper/utils"
)

// Install install installer
func Install(staging, force bool) error {
	if isInstallerExist() && !force {
		fmt.Println("SDKBox has been installed")
		return nil
	}
	sdkboxHome := env.SDKBoxHome()
	tempPath, err := utils.CURL(url.GetInstallerURL(staging), filepath.Join(sdkboxHome, "bin", "sdkbox_installer.zip"))
	if nil != err {
		panic(err)
	}

	if err := utils.Unzip(tempPath, filepath.Join(sdkboxHome, "bin")); nil != err {
		fmt.Println("Unzip sdkbox installer failed")
		panic(err)
	}

	fmt.Println("")
	fmt.Println(">>>")
	fmt.Println("Please add follow to your environment path:")
	fmt.Println("export SDKBOX_HOME=" + sdkboxHome)
	fmt.Println("export PATH=${SDKBOX_HOME}/bin:$PATH")
	fmt.Println(">>>")
	fmt.Println("")
	return nil
}

func isInstallerExist() bool {
	sdkboxBin := filepath.Join(env.SDKBoxHome(), "bin", "sdkbox")
	return utils.Exist(sdkboxBin)
}
