package install

import (
	"github.com/intelfike/mulpage/types"
)

type Package struct{}

func (p *Package) Define(pack *types.Package) {
	pack.Init("package", "インストール")

	pack.SetMethod("Index", "インストール", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})
}
