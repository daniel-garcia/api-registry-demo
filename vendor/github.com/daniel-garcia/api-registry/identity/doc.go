package identity

import (
	"embed"

	"github.com/daniel-garcia/api-registry"
)



//go:embed infoblox/*
//go:embed *.yaml
var content embed.FS

func init() {
	apiregistry.Register("identity", content)
}

