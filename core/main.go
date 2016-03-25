package core

import ()

//go:generate -command coremodels go run ../generators/core.go
//go:generate coremodels

type Model interface {
	GetId() ModelId
}

type Device struct {
	Models []Model
}
