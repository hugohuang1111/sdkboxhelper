package installer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"sdkboxhelper/env"
	"sdkboxhelper/url"
	"sdkboxhelper/utils"

	"github.com/hashicorp/go-version"
	cp "github.com/otiai10/copy"
)

// CreatorProjectInfo creator project info
type CreatorProjectInfo struct {
	Engine   string `json:"engine"`
	Version  string `json:"version"`
	Packages string `json:"packages"`
}

// InstallCreatorPlugin install sdkbox gui for creator
func InstallCreatorPlugin(staging, force bool, project string) error {
	cpi := loadCreatorProjectInfo(project)

	if "" == cpi.Version {
		return errors.New("wrong creator project path")
	}

	curVer, err := version.NewVersion(cpi.Version)
	if err != nil {
		return err
	}
	targetVer, err := version.NewVersion("2.4.1")
	if err != nil {
		return err
	}
	if curVer.LessThan(targetVer) {
		fmt.Printf("Cocos Creator(%v) has built-in SDKBox, needn't install", cpi.Version)
		fmt.Println("")
		return nil
	}

	packagesDir := filepath.Join(project, cpi.Packages)
	packagePath := filepath.Join(env.UserHome(), ".CocosCreator", "packages", "sdkbox")
	packageExist := utils.Exist(packagePath)
	if !packageExist {
		packagePath = filepath.Join(packagesDir, "sdkbox")
		packageExist = utils.Exist(packagePath)
	}
	if packageExist {
		fmt.Println("SDKBox plugin have been installed")
		return nil
	}

	sdkboxHome := env.SDKBoxHome()
	tempPath, err := utils.CURL(url.GetInstallerCreatorGUIURL(staging), filepath.Join(sdkboxHome, "temp", "sdkboxguiforcreator.zip"))
	if nil != err {
		panic(err)
	}
	if err := utils.Unzip(tempPath, filepath.Join(sdkboxHome, "temp")); nil != err {
		fmt.Println("Unzip sdkbox creator plugin failed")
		panic(err)
	}

	cp.Copy(filepath.Join(sdkboxHome, "temp", "sdkbox"), filepath.Join(packagesDir, "sdkbox"))
	utils.MakeSureDirExist(filepath.Join(sdkboxHome, "creator", "app"))
	os.Rename(filepath.Join(sdkboxHome, "temp", "sdkbox", "app"), filepath.Join(sdkboxHome, "creator", "app"))

	fmt.Println("")
	fmt.Println(">>>")
	fmt.Println("SDKBox GUI installed")
	fmt.Println(">>>")
	fmt.Println("")

	return nil
}

func loadCreatorProjectInfo(project string) CreatorProjectInfo {
	cpi := CreatorProjectInfo{}
	projectJSON := filepath.Join(project, "project.json")
	if !utils.Exist(projectJSON) {
		return cpi
	}

	content, err := ioutil.ReadFile(projectJSON)
	if err != nil {
		return cpi
	}
	err = json.Unmarshal(content, &cpi)
	if err != nil {
		return cpi
	}

	return cpi
}

func isSDKBoxCreatorGUIExist(packagesDir, plugin string) bool {
	sdkboxPlugin := filepath.Join(packagesDir, plugin)
	return utils.Exist(sdkboxPlugin)
}
