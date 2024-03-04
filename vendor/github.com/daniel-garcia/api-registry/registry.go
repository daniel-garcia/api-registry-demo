package apiregistry

import (
	"embed"
	"fmt"
)

var packages = make(map[string]embed.FS)


func Register(name string, fs embed.FS) {
	_, found := packages[name]
	if found {
		panic(fmt.Sprintf("package %v already registered", name))
	}
	packages[name] = fs
}

func Packages() map[string]embed.FS {
	return packages
}

