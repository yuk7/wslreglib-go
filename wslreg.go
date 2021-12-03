package wslreg

import (
	"errors"
	"io"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

const (
	// LxssBaseRoot is CURRENT_USER
	LxssBaseRoot = registry.CURRENT_USER
	// LxssBaseRootStr is CURRENT_USER string
	LxssBaseRootStr = "HKEY_CURRENT_USER"
	// LxssBaseKey is path of lxss registry
	LxssBaseKey = "Software\\Microsoft\\Windows\\CurrentVersion\\Lxss"
	// WsldlTermKey is registry key name used for wsldl terminal infomation
	WsldlTermKey = "wsldl-term"
	// FlagWsldlTermDefault is default terminal (conhost)
	FlagWsldlTermDefault = 0
	// FlagWsldlTermWT is Windows Terminal
	FlagWsldlTermWT = 1
	// FlagWsldlTermFlute is Fluent Terminal
	FlagWsldlTermFlute = 2
	// InvalidNum is Num used for invalid
	InvalidNum = -1
)

// GetLxUuidList gets guid key lists
func GetLxUuidList() (uuidList []string, err error) {
	baseKey, tmpErr := registry.OpenKey(LxssBaseRoot, LxssBaseKey, registry.READ)
	if tmpErr != nil && tmpErr != io.EOF {
		err = tmpErr
		return
	}
	uuidList, tmpErr = baseKey.ReadSubKeyNames(1024)
	if tmpErr != nil && tmpErr != io.EOF {
		err = tmpErr
		return
	}
	return
}

// GetProfileFromName gets distro profile from name
func GetProfileFromName(distributionName string) (profile Profile, err error) {
	uuidList, tmpErr := GetLxUuidList()
	if tmpErr != nil {
		err = tmpErr
		return
	}

	errStr := ""
	for _, loopUUID := range uuidList {
		profile, _ = ReadProfile(loopUUID)
		if strings.EqualFold(profile.DistributionName, distributionName) {
			return
		}
	}
	err = errors.New("Registry Key Not found\n" + errStr)
	profile = NewProfile()
	return
}

// GetProfileFromBasePath gets distro profile from BasePath
func GetProfileFromBasePath(basePath string) (profile Profile, err error) {
	uuidList, tmpErr := GetLxUuidList()
	if tmpErr != nil {
		err = tmpErr
		return
	}

	basePathAbs, tmpErr := filepath.Abs(basePath)
	if err != nil {
		basePathAbs = basePath
	}

	errStr := ""
	for _, loopUUID := range uuidList {
		profile, _ = ReadProfile(loopUUID)
		if strings.EqualFold(profile.BasePath, basePathAbs) {
			return
		}
	}
	err = errors.New("Registry Key Not found\n" + errStr)
	profile = NewProfile()
	return
}
