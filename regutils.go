package wslreg

import "golang.org/x/sys/windows/registry"

// regSetStringWithCmdAndFix sets value to registry key and do fix with reg command
func regSetStringWithCmdAndFix(regkey registry.Key, regpathStr, keyname, value string) error {
	// backup oldValue
	oldValue := ""
	if regkey != 0 {
		oldValue, _, _ = regkey.GetStringValue(value)
	}

	// write with external command
	err := regSetStringWithCmd(regpathStr, keyname, value)
	if err != nil {
		return err
	}

	// if regkey not 0, check if the value was written correctly
	if regkey != 0 && oldValue != "" && oldValue != value {
		newValue, _, _ := regkey.GetStringValue(value)
		if oldValue == newValue {
			// not changed, maybe appx virtual registry used
			// delete it and rewrite registry value
			regkey.DeleteValue(keyname)
			err := regSetStringWithCmd(regpathStr, keyname, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// regSetDWordWithCmdAndFix sets value to registry key and do fix with reg command
func regSetDWordWithCmdAndFix(regkey registry.Key, regpathStr, keyname string, value uint32) error {
	// backup oldValue
	oldValue := InvalidNum
	if regkey != 0 {
		val, _, _ := regkey.GetIntegerValue(keyname)
		oldValue = int(val)
	}

	// write with external command
	err := regSetDWordWithCmd(regpathStr, keyname, value)
	if err != nil {
		return err
	}

	// if regkey not 0, check if the value was written correctly
	if regkey != 0 && oldValue != InvalidNum && oldValue != int(value) {
		newValue, _, _ := regkey.GetIntegerValue(keyname)
		if oldValue == int(newValue) {
			// not changed, maybe appx virtual registry used
			// delete it and rewrite registry value
			regkey.DeleteValue(keyname)
			err := regSetDWordWithCmd(regpathStr, keyname, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
