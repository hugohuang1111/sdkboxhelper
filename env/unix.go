package env

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func addUnixEnvVar(key string, value string) error {
	profilePath := getUnixProfileFile()
	if "" == profilePath {
		return errors.New("profile file is empty")
	}

	s1 := fmt.Sprintf("export %s=%s", key, value)
	s2 := fmt.Sprintf("export PATH=${%s}/bin:$PATH", key)

	ioutil.WriteFile(profilePath, []byte(s1), os.ModeAppend)
	ioutil.WriteFile(profilePath, []byte(s2), os.ModeAppend)

	return nil
}

func getUnixProfileFile() string {
	userHome := UserHome()
	switch runtime.GOOS {
	case "windows":
		// import "golang.org/x/sys/windows/registry"
	case "linux", "darwin":
		profile := ""
		if isZsh() {
			profile = ".zshrc"
		} else {
			if "linux" == runtime.GOOS {
				profile = ".bashrc"
			} else if "darwin" == runtime.GOOS {
				profile = ".bash_profile" // ".bash_login", ".profile"
			}
		}
		return filepath.Join(userHome, profile)
	default:
		fmt.Println(runtime.GOOS)
	}

	return ""
}

func isZsh() bool {
	val := os.Getenv("SHELL")
	return strings.HasSuffix(val, "zsh")
}
