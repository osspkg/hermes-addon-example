package main

import (
	hermesaddons "github.com/osspkg/hermes-addons"
)

func (Example) PkgName() string {
	return "com.osspkg.hermes-addon-example"
}

func (Example) Version() hermesaddons.SemVersion {
	return "v0.0.1"
}

func (Example) GetIcon(size hermesaddons.IconSize) string {
	return "https://avatars.githubusercontent.com/u/134978195?s=200&v=4"
}

func (Example) Info() hermesaddons.Info {
	return hermesaddons.Info{
		Name:        "Example App",
		Author:      "OSSPkg Team",
		Email:       "hermes@osspkg.com",
		Title:       "Example application",
		Description: "",
		Info:        "",
	}
}
