package support

import (
	"github.com/intelfike/mulpage/types"
)

type Package struct{}

func (p *Package) Define(pack *types.Package) {
	pack.Init("package", "サポート情報")

	pack.SetMethod("Index", "サポート情報", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})
}
