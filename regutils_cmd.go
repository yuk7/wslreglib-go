package wslreg

import (
	"os/exec"
	"strconv"
)

// regSetStringWithCmd writes DWord key with command, forcibly use the real registry
func regSetStringWithCmd(regpath, keyname, value string) error {
	regexe := getWindowsDirectory() + "\\System32\\reg.exe"

	_, err := exec.Command(regexe, "add", regpath, "/v", keyname, "/t", "REG_SZ", "/d", value, "/f").Output()
	return err
}

// regSetDWordWithCmd writes DWord key with command, forcibly use the real registry
func regSetDWordWithCmd(regpath, keyname string, value uint32) error {
	regexe := getWindowsDirectory() + "\\System32\\reg.exe"

	_, err := exec.Command(regexe, "add", regpath, "/v", keyname, "/t", "REG_DWORD", "/d", strconv.Itoa(int(value)), "/f").Output()
	return err
}

// SetWslVersion sets wsl version
func SetWslVersion(distributionName string, version int) error {
	wslexe := getWindowsDirectory() + "\\System32\\wsl.exe"
	_, err := exec.Command(wslexe, "--set-version", distributionName, strconv.Itoa(version)).Output()
	return err
}
