package main

import (
	"context"
	"encoding/json"

	hermesaddons "github.com/osspkg/hermes-addons"
)

func (v Example) ACL() hermesaddons.ACLGetter {
	//TODO implement me
	panic("implement me")
}

func (v Example) Form(id uint) (json.Marshaler, error) {
	//TODO implement me
	panic("implement me")
}

func (v Example) Call(ctx context.Context, id uint, data []byte, user hermesaddons.UserGetter) (json.Marshaler, error) {
	//TODO implement me
	panic("implement me")
}

func (v *Example) PipeCall(ctx context.Context, id uint, data []byte) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}
