package top

import (
	"github.com/intelfike/mulpage/types"
)

type Package struct{}

func (p *Package) Define(pack *types.Package) {
	pack.Init("package", "トップページ")

	pack.Before = func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.Assign("Title", "トップページ")
		return nil
	}

	pack.SetMethod("Index", "トップページ", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})
}
