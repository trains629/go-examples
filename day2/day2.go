package day2

import (
	reg "golang.org/x/sys/windows/registry"
)

func ReadStringKey(rootKey reg.Key, pathKey string, key string) (string, error) {
	k, err := reg.OpenKey(rootKey, pathKey, reg.ALL_ACCESS)
	if err != nil {
		return "", err
	}
	defer k.Close()

	s, _, err := k.GetStringValue(key)
	if err != nil {
		return "", err
	}
	return s, nil
}
