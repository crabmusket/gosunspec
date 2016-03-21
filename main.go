package gosunspec

import ()

//go:generate -command models go run generators/models.go
//go:generate models

type Model interface {
	GetId() ModelId
}
