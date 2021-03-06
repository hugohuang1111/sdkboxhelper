package utils

import (
	"io/ioutil"
	"net/http"
	"os"
)

// CURL curl
func CURL(url string, path string) (s string, e error) {
	resp, err := http.Get(url)
	if nil != err {
		return "", err
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	if "" != path {
		MakeSureDirExist(path)
		err = ioutil.WriteFile(path, content, os.FileMode(0666))
		return path, err
	}

	return string(content), nil
}
