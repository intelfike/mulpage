package support

import (
	"github.com/intelfike/mulpage/types"
)

type Package struct{}

func (p *Package) Define(pack *types.Package) {
	pack.Init("サポート情報")

	pack.SetMethod("Index", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.Assign("Title", "サポート情報")
		return nil
	})
}
