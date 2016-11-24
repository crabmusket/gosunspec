package core

type Model interface {
	GetId() ModelId
}

type Device struct {
	Models []Model
}
