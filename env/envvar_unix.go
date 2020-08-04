// +build !windows

package env

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sdkboxhelper/utils"
	"strings"
)

func addEnvVarImp(key string, value string) error {
	profilePath := getUnixProfileFile()
	if "" == profilePath {
		return errors.New("profile file is empty")
	}

	s1 := fmt.Sprintf("\nexport %s=%s\nexport PATH=${%s}/bin:$PATH\n", key, value, key)

	utils.AppendToFile(profilePath, []byte(s1))

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
