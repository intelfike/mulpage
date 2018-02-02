package top

import (
	"github.com/intelfike/mulpage/types"
)

type Package struct{}

func (p *Package) Define(pack *types.Package) {
	pack.Init("package", "top")

	pack.SetMethod("Index", "top", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})
}
