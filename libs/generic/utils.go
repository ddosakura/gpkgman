package generic

import (
	"os"
	"path"

	"github.com/ddosakura/gklang"
)

func getDepDir() string {
	dep := path.Join(".", "dep")
	if err := os.MkdirAll(dep, 0755); err != nil {
		gklang.Er(err)
	}
	return dep
}
