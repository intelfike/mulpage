package usage

import (
	"github.com/intelfike/mulpage/types"
)

type Package struct{}

func (p *Package) Title() string {
	return "isearの使い方"
}

func (p *Package) Define(pack *types.Package) {
	pack.SetMethod("Index", func(info *types.PageInfo) *types.Redirect {
		info.Assign("Title", "isearの使い方")
		return nil
	})
}
