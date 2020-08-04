package main

import (
	"flag"
	"fmt"

	"sdkboxhelper/installer"
	"sdkboxhelper/utils"
)

const sdkboxHelperVersion string = "0.0.5"

// params
var cmd string
var itype string
var staging bool
var projectPath string
var force bool

// http://staging.sdkbox.com/gui/creator/sdkbox-1.0.4.zip

func init() {
	// flag.StringVar(&cmd, "c", "install", "command")
	flag.StringVar(&itype, "t", "installer", "install type, installer or creator, default is installer")
	flag.StringVar(&projectPath, "p", "", "project path")
	flag.BoolVar(&staging, "staging", false, "staging server")
	flag.BoolVar(&force, "force", false, "force upgrade")
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
	flag.Parse()

	if "" == projectPath {
		projectPath = utils.CurDir()
	}
	fmt.Println("SDKBox Helper Version:" + sdkboxHelperVersion)
	fmt.Println("")

	switch itype {
	case "installer":
		installer.Install(staging, force)
	case "creator":
		installer.InstallCreatorPlugin(staging, force, projectPath)
	default:
		panic("unknow command type, should be installer or creator")
	}

	// switch cmd {
	// case "install":
	// case "upgrade":
	// default:
	// 	panic("unknow command")
	// }

	// fmt.Println(getInstallerVersionURL())

	fmt.Println(">>> Done")
	fmt.Println("")
}
