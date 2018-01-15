package top

import (
	"github.com/intelfike/mulpage/dev/global"
	"github.com/intelfike/mulpage/dev/types"
)

type Package struct{}

func (p *Package) Init(pack types.Package) {
	pack.SetMethod("Index", func() *types.Redirect {
		info := global.PageInfo
		info.Title = "isearweb"
		info.Assign("mod", "modmod")
		return nil
	})

	pack.SetMethod("New", func() *types.Redirect {
		info := global.PageInfo
		info.Template = "Index"
		info.Assign("mod", "new")
		return nil
	})
}
