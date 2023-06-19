package main

import (
	"context"

	"github.com/osspkg/hermes-addons/api1"
	"github.com/osspkg/hermes-addons/base"
)

func HermesAPI() api1.Api {
	return &Example{}
}

type Example struct {
}

func (v *Example) Inject(dic base.DIContainer) error {
	return nil
}

func (v *Example) Dependency() []string {
	return []string{
		"com.osspkg.database",
	}
}

func (v *Example) Up(ctx context.Context) error {
	return nil
}

func (v *Example) Down() error {
	return nil
}
