package install

import (
	"github.com/intelfike/mulpage/types"
)

type Package struct{}

func (p *Package) Define(pack *types.Package) {
	pack.Init("インストール")

	pack.SetMethod("Index", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.Assign("Title", "インストール")
		return nil
	})
}
