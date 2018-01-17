package top

import (
	"github.com/intelfike/mulpage/types"
)

type Package struct{}

func (p *Package) Define(pack *types.Package) {
	pack.Init("トップページ")

	pack.Before = func(info *types.PageInfo) *types.Redirect {
		info.Assign("Title", "トップページ")
		return nil
	}

	pack.SetMethod("Index", func(info *types.PageInfo) *types.Redirect {
		info.Assign("mod", "Index")
		return nil
	})

	pack.SetMethod("New", func(info *types.PageInfo) *types.Redirect {
		info.Template = "Index"
		info.Assign("mod", "New")
		return nil
	})
}
