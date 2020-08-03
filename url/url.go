package url

func GetInstallerURL() string {
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
