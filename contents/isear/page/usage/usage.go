package usage

import (
	"github.com/intelfike/mulpage/types"
)

type Package struct{}

func (p *Package) Define(pack *types.Package) {
	pack.Init("isearの使い方")

	pack.SetMethod("Index", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.Assign("Title", "isearの使い方")
		return nil
	})
}
