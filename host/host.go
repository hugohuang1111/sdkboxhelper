package host

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type location struct {
	CountryCode string `json:"countryCode"`
}

func getHost(staging bool) string {
	if staging {
		return "http://staging.sdkbox.com/"
	}
	countryCode := loadLocation()
	if "CN" == countryCode {
		return "http://sdkbox.anysdk.com/"
	}
	return "http://download.sdkbox.com/"
}

func loadLocation() string {
	// ~/.sdkbox/conf/loc.json
	locPath := filepath.Join(sdkboxHome, "conf", "loc.json")
	locBytes, err := ioutil.ReadFile(locPath)
	loc := location{}
	err = json.Unmarshal(locBytes, &loc)
	if nil != err || "" == loc.CountryCode {
		loc.CountryCode = loadLocationByIP()
		if locBytes, err = json.Marshal(loc); nil == err {
			ioutil.WriteFile(locPath, locBytes, os.FileMode(0666))
		}
	}
	return loc.CountryCode
}

func loadLocationByIP() string {
	// "http://myexternalip.com/raw"
	// "http://ip-api.com/json/"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://ip-api.com/json/", nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.3")
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)

	var locMap map[string]interface{}
	if err := json.Unmarshal(content, &locMap); err != nil {
		return ""
	}

	return locMap["countryCode"].(string)
}
