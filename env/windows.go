package env

import (
	"fmt"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func addWinEnvVar(key string, value string) error {

	k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.QUERY_VALUE | registry.WRITE)
	if err != nil {
		return err
	}
	defer k.Close()

	err = k.SetStringValue(key, value)
	if err != nil {
		return err
	}

	pathKey := "Path"
	s, _, err := k.GetStringValue(pathKey)
	if err != nil {
		return err
	}
	if !strings.Contains(s, key) {
		s = fmt.Sprintf("%s;%%%s%%\\bin", s, key)
		k.SetStringValue(pathKey, s)
	}

	return nil
}
