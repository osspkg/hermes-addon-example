package main

import (
	"github.com/osspkg/hermes-addons/api1"
	"github.com/osspkg/hermes-addons/base"
)

func (Example) Schema() uint {
	return 1
}

func (Example) PkgName() string {
	return "com.osspkg.hermes-addon-example"
}

func (Example) Version() base.SemVersion {
	return "v0.0.1"
}

func (Example) GetIcon(size base.IconSize) string {
	return "https://avatars.githubusercontent.com/u/134978195?s=200&v=4"
}

func (Example) Info() api1.Info {
	return api1.Info{
		Name:        "Example App",
		Author:      "OSSPkg Team",
		Email:       "hermes@osspkg.com",
		Title:       "Example application",
		Description: "",
		Info:        "",
	}
}
