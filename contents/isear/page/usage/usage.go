package usage

import (
	"github.com/intelfike/mulpage/types"
)

type Package struct{}

func (p *Package) Define(pack *types.Package) {
	pack.Init("package", "isearの使い方")

	pack.Before = func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		if tpl.Template == "Index" {
			tpl.Template = "Functions"
		}
		return nil
	}

	pack.SetMethod("Functions", "機能", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})

	pack.SetMethod("Name", "名称・用語", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})

	pack.SetMethod("SKey", "ショートカットキー", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})
}
