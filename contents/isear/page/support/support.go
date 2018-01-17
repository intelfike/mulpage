package support

import (
	"github.com/intelfike/mulpage/types"
)

type Package struct{}

func (p *Package) Title() string {
	return "サポート情報"
}

func (p *Package) Define(pack *types.Package) {
	pack.SetMethod("Index", func(info *types.PageInfo) *types.Redirect {
		info.Assign("Title", "サポート情報")
		return nil
	})
}
