package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"sdkbox.com/helper/utils"
)

// params
var cmd string
var itype string
var staging bool

// http://staging.sdkbox.com/gui/creator/sdkbox-1.0.4.zip

func init() {
	flag.StringVar(&cmd, "c", "install", "command")
	flag.StringVar(&itype, "t", "installer", "install type")
	flag.BoolVar(&staging, "staging", false, "staging server")
}

func installInstaller() {
	sdkboxHome = "" // env.GetSDKBoxHome()
	tempPath, err := curl(getInstallerURL(), filepath.Join(sdkboxHome, "bin", "sdkbox_installer.zip"))
	if nil != err {
		panic(err)
	}
	if err := utils.Unzip(tempPath, filepath.Join(sdkboxHome, "bin")); nil != err {
		fmt.Println("Unzip sdkbox installer failed")
		panic(err)
	}

	fmt.Println("")
	fmt.Println(">>>")
	fmt.Println("Please add follow to your environment path")
	fmt.Println("")
	fmt.Println("export SDKBOX_HOME=" + sdkboxHome)
	fmt.Println("export PATH=${SDKBOX_HOME}/bin:$PATH")
	fmt.Println(">>>")
	fmt.Println("")
}



/*
 * useage:
 * sdkboxhelper == sdkboxhelper install installer
 * sdkboxhelper install gui_for_creator
 * sdkboxhelper upgrade installer
 * sdkboxhelper upgrade gui_for_creator
 *
 */
func main() {
	fmt.Println(">>> SDKBox Entry")
	flag.Parse()

	switch cmd {
	case "install":
		switch itype {
		case "installer":
			installInstaller()
		case "gui_for_creator":
			// url := getInstallerCreatorGUIVersionURL()
			// fmt.Println(url)
		default:
			panic("unknow command type")
		}
	case "upgrade":
	default:
		panic("unknow command")
	}

	// fmt.Println(getInstallerVersionURL())

	fmt.Println(">>> Done")
}
