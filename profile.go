package wslreg

import (
	"errors"
	"io"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/sys/windows/registry"
)

// NewProfile creates empty profile
func NewProfile() Profile {
	profile := Profile{}
	profile.DefaultUid = InvalidNum
	profile.Flags = InvalidNum
	profile.State = InvalidNum
	profile.Version = InvalidNum
	profile.WsldlTerm = InvalidNum
	return profile
}

// GenerateProfile generates new profile with UUID
func GenerateProfile() Profile {
	profile := NewProfile()
	profile.UUID = "{" + uuid.NewV4().String() + "}"
	profile.DefaultUid = 0x0
	profile.Flags = 0x7
	profile.State = 0x1
	profile.Version = 0x2
	profile.WsldlTerm = 0x0
	return profile
}

// WriteProfile writes profile to registry
func WriteProfile(profile Profile) error {
	if profile.UUID == "" {
		return errors.New("Empty UUID")
	}
	key, _, err := registry.CreateKey(LxssBaseRoot, LxssBaseKey+"\\"+profile.UUID, registry.ALL_ACCESS)
	if err != nil {
		key = 0
	}
	regpathStr := LxssBaseRootStr + "\\" + LxssBaseKey + "\\" + profile.UUID

	if profile.BasePath != "" {
		err = regSetStringWithCmdAndFix(key, regpathStr, "BasePath", profile.BasePath)
		if err != nil {
			return err
		}
	}
	if profile.DistributionName != "" {
		err = regSetStringWithCmdAndFix(key, regpathStr, "DistributionName", profile.DistributionName)
		if err != nil {
			return err
		}
	}
	if profile.DefaultUid != InvalidNum {
		err = regSetDWordWithCmdAndFix(key, regpathStr, "DefaultUid", uint32(profile.DefaultUid))
		if err != nil {
			return err
		}
	}

	if profile.Flags != InvalidNum {
		err = regSetDWordWithCmdAndFix(key, regpathStr, "Flags", uint32(profile.Flags))
		if err != nil {
			return err
		}
	}

	if profile.State != InvalidNum {
		err = regSetDWordWithCmdAndFix(key, regpathStr, "State", uint32(profile.State))
		if err != nil {
			return err
		}
	}

	if profile.Version != InvalidNum {
		err = regSetDWordWithCmdAndFix(key, regpathStr, "Version", uint32(profile.Version))
		if err != nil {
			return err
		}
	}

	if profile.PackageFamilyName != "" {
		err = regSetStringWithCmdAndFix(key, regpathStr, "PackageFamilyName", profile.PackageFamilyName)
		if err != nil {
			return err
		}
	}
	if profile.WsldlTerm != InvalidNum {
		err = regSetDWordWithCmdAndFix(key, regpathStr, WsldlTermKey, uint32(profile.WsldlTerm))
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadProfile reads profile from registry
func ReadProfile(lxUuid string) (profile Profile, err error) {
	profile = NewProfile()
	profile.UUID = lxUuid
	key, err := registry.OpenKey(LxssBaseRoot, LxssBaseKey+"\\"+profile.UUID, registry.READ)
	if err != nil {
		return
	}
	basepath, _, tmperr := key.GetStringValue("BasePath")
	if tmperr == nil || tmperr == io.EOF {
		profile.BasePath = basepath
	}
	distributionName, _, tmperr := key.GetStringValue("DistributionName")
	if tmperr == nil || tmperr == io.EOF {
		profile.DistributionName = distributionName
	}
	flags, _, tmperr := key.GetIntegerValue("Flags")
	if tmperr == nil || tmperr == io.EOF {
		profile.Flags = int(flags)
	}
	state, _, tmperr := key.GetIntegerValue("State")
	if tmperr == nil || tmperr == io.EOF {
		profile.State = int(state)
	}
	version, _, tmperr := key.GetIntegerValue("Version")
	if tmperr == nil || tmperr == io.EOF {
		profile.Version = int(version)
	}
	wsldlTerm, _, tmperr := key.GetIntegerValue(WsldlTermKey)
	if tmperr == nil || tmperr == io.EOF {
		profile.WsldlTerm = int(wsldlTerm)
	}
	pkgName, _, tmperr := key.GetStringValue("PackageFamilyName")
	if tmperr == nil || tmperr == io.EOF {
		profile.PackageFamilyName = pkgName
	}
	return
}
