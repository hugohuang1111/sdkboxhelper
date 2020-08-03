package url

import (
	"encoding/json"

	"sdkbox.com/helper/utils"
)

// GetInstallerURL get installer url
func GetInstallerURL(staging bool) string {
	url := getInstallerVersionURL(staging)
	content, err := utils.CURL(url, "")
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
		valMap := value.(map[string]interface{})
		return getHost(staging) + "installer/v1/" + valMap["bundle"].(string)
	}

	return ""
}

// GetInstallerCreatorGUIURL get installer creator gui url
func GetInstallerCreatorGUIURL(staging bool) string {
	url := getInstallerCreatorGUIVersionURL(staging)
	content, err := utils.CURL(url, "")
	if nil != err {
		panic(err)
	}
	var versionInfo map[string]interface{}
	if err := json.Unmarshal([]byte(content), &versionInfo); err != nil {
		return ""
	}

	// http://sdkbox.anysdk.com/gui/creator/version
	versionStr := versionInfo["version"].(string)

	// http://staging.sdkbox.com/gui/creator/sdkbox-1.0.4.zip
	return getHost(staging) + "gui/creator/sdkbox-" + versionStr + ".zip"
}

func getInstallerCreatorGUIVersionURL(staging bool) string {
	return getHost(staging) + "gui/creator/version"
}

func getInstallerVersionURL(staging bool) string {
	return getHost(staging) + "installer/v1/manifest.json"
}
