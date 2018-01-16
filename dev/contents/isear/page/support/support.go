package support

import (
	"github.com/intelfike/mulpage/dev/ifc"
	"github.com/intelfike/mulpage/dev/types"
)

type Package struct{}

var _ ifc.Package = &Package{}

func (p *Package) Init(pack *types.Package) {
	pack.SetMethod("Index", func(info *types.PageInfo) *types.Redirect {
		info.Assign("Title", "サポート情報")
		return nil
	})
}
