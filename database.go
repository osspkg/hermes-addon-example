package main

import (
	hermesaddons "github.com/osspkg/hermes-addons"
)

func (Example) Database() []hermesaddons.DatabaseMigration {
	return []hermesaddons.DatabaseMigration{
		{
			ID:   "0001",
			Up:   "",
			Down: "",
		},
	}
}
