package main

import "github.com/osspkg/hermes-addons/base"

func (Example) Database() []base.DatabaseMigration {
	return []base.DatabaseMigration{
		{
			ID:   "0001",
			Up:   "",
			Down: "",
		},
	}
}
