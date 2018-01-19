package top

import (
	"github.com/intelfike/mulpage/types"
)

type Package struct{}

func (p *Package) Define(pack *types.Package) {
	pack.Init("トップページ")

	pack.Before = func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.Assign("Title", "トップページ")
		return nil
	}

	pack.SetMethod("Index", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.Assign("mod", "Index")
		return nil
	})

	pack.SetMethod("New", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.Template = "Index"
		tpl.Assign("mod", "New")
		return nil
	})
}
