package main

import "github.com/osspkg/hermes-addons/base"

func (Example) Database() base.MigrationsGetter {
	return &Database{}
}

type Database struct{}

func (Database) Table() string {
	return "example_table"
}

func (Database) Data() []base.DatabaseMigration {
	return []base.DatabaseMigration{
		{
			ID:   "0001",
			Up:   "",
			Down: "",
		},
	}
}
