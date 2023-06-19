package main

import (
	"context"

	hermesaddons "github.com/osspkg/hermes-addons"
	"github.com/osspkg/hermes-addons/dependency"
)

func HermesAPI() hermesaddons.Api {
	return &Example{}
}

type Example struct {
}

func (v *Example) Inject(dic hermesaddons.DIContainer) error {
	return nil
}

func (v *Example) Dependency() []string {
	return []string{
		dependency.Database,
	}
}

func (v *Example) Up(ctx context.Context) error {
	return nil
}

func (v *Example) Down() error {
	return nil
}
