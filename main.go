package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"path/filepath"

	"sdkbox.com/helper/utils"
)

var sdkboxHome string

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

func getInstallerURL() string {
	url := getInstallerVersionURL()
	content, err := curl(url, "")
	if nil != err {
		panic(err)
	}
	var manifestMap map[string]interface{}
	if err := json.Unmarshal([]byte(content), &manifestMap); err != nil {
		return ""
	}

	item := manifestMap["packages"].(map[string]interface{})
	item = item["SDKBOX"].(map[string]interface{})
	item = item["versions"].(map[string]interface{})
	for _, value := range item {
		valMap := value.(map[string]string)
		return getHost() + "installer/v1/" + valMap["bundle"]
	}

	return ""
}

func getInstallerCreatorGUIVersionURL() string {
	return getHost() + "gui/creator/version"
}

func getInstallerVersionURL() string {
	return getHost() + "installer/v1/manifest.json"
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

	userHome, err := utils.Home()
	if nil != err {
		panic(err)
	}
	sdkboxHome = filepath.Join(userHome, ".sdkbox")

	switch cmd {
	case "install":
		switch itype {
		case "installer":
			installInstaller()
		case "gui_for_creator":
			url := getInstallerCreatorGUIVersionURL()
			fmt.Println(url)
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
