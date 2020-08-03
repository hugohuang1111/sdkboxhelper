package main

import (
	"flag"
	"fmt"

	"sdkbox.com/helper/installer"
	"sdkbox.com/helper/utils"
)

// params
var cmd string
var itype string
var staging bool
var projectPath string

// http://staging.sdkbox.com/gui/creator/sdkbox-1.0.4.zip

func init() {
	flag.StringVar(&cmd, "c", "install", "command")
	flag.StringVar(&itype, "t", "installer", "install type")
	flag.StringVar(&projectPath, "p", "", "project path")
	flag.BoolVar(&staging, "staging", false, "staging server")
}

/*
 * useage:
 * sdkboxhelper == sdkboxhelper install installer
 * sdkboxhelper install creator
 * sdkboxhelper upgrade installer
 * sdkboxhelper upgrade creator
 *
 */
func main() {
	fmt.Println(">>> SDKBox Entry")
	flag.Parse()

	if "" == projectPath {
		projectPath = utils.CurDir()
	}

	switch cmd {
	case "install":
		switch itype {
		case "installer":
			installer.Install(staging, false)
		case "creator":
			installer.InstallCreatorPlugin(staging, false, projectPath)
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
