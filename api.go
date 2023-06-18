package main

import (
	"context"
	"encoding/json"

	"github.com/osspkg/hermes-addons/api1"
)

func (v Example) ACL() api1.ACLGetter {
	//TODO implement me
	panic("implement me")
}

func (v Example) Form(id uint) (json.Marshaler, error) {
	//TODO implement me
	panic("implement me")
}

func (v Example) Call(ctx context.Context, id uint, data []byte, user api1.UserGetter) (json.Marshaler, error) {
	//TODO implement me
	panic("implement me")
}
