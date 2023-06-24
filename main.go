package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/osspkg/hermes-addons/dependency"

	hermesaddons "github.com/osspkg/hermes-addons"
)

func HermesAPI() hermesaddons.Api {
	return &Example{}
}

type Example struct {
	orm dependency.ORM
}

func (e *Example) Inject(dic hermesaddons.DIContainer) error {
	orm, err := dic.Get(dependency.Database)
	if err != nil {
		return err
	}
	ormCast, ok := orm.(dependency.ORM)
	if !ok {
		return fmt.Errorf("cast dependency `%s` to `dependency.ORM`", dependency.Database)
	}
	e.orm = ormCast

	return nil
}

func (e *Example) Up(ctx context.Context) error {
	return nil
}

func (e *Example) Down() error {
	return nil
}

func (e *Example) Database() []hermesaddons.DatabaseMigration {
	//TODO implement me
	panic("implement me")
}

func (e *Example) Form(id uint) (json.Marshaler, error) {
	//TODO implement me
	panic("implement me")
}

func (e *Example) Call(ctx context.Context, id uint, data []byte, user hermesaddons.UserGetter) (json.Marshaler, error) {
	//TODO implement me
	panic("implement me")
}

func (e *Example) ACL() []hermesaddons.ACLModel {
	//TODO implement me
	panic("implement me")
}
