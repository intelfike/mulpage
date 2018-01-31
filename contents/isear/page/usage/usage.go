package usage

import (
	"github.com/intelfike/mulpage/types"
)

type Package struct{}

func (p *Package) Define(pack *types.Package) {
	pack.Init("isearの使い方")

	pack.SetMethod("Index", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.Assign("Title", "目次")
		return nil
	})

	pack.SetMethod("Functions", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})

	pack.SetMethod("Name", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.Assign("Title", "名称・用語")
		return nil
	})

	pack.SetMethod("SKey", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.Assign("Title", "ショートカットキー")
		return nil
	})
}
