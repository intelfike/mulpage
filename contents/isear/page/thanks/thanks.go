package thanks

import (
	"github.com/intelfike/mulpage/types"
)

type Package struct{}

func (p *Package) Define(pack *types.Package) {
	pack.Init("package", "thanks")

	pack.SetMethod("Index", "Thanks", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})
}
