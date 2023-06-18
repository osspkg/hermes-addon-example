package main

import (
	"context"

	"github.com/osspkg/hermes-addons/api1"
	"github.com/osspkg/hermes-addons/base"
)

var HermesAPI api1.Api = &Example{}

type Example struct {
}

func (v *Example) Inject(dic base.DIContainer) error {
	return nil
}

func (v *Example) Dependency() []string {
	return nil
}

func (v *Example) Up(ctx context.Context) error {
	return nil
}

func (v *Example) Down() error {
	return nil
}
