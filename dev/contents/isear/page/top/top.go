package top

import (
	"fmt"

	"github.com/intelfike/mulpage/dev/ifc"
	"github.com/intelfike/mulpage/dev/types"
)

type Package struct{}

var _ ifc.Package = &Package{}

func (p *Package) Init(pack types.Package) {
	pack.Before = func(info *types.PageInfo) *types.Redirect {
		fmt.Println("hello!!!")
		return nil
	}
	pack.SetMethod("Index", func(info *types.PageInfo) *types.Redirect {
		info.Assign("Title", "トップページ")
		info.Assign("mod", "modmod")
		return nil
	})

	pack.SetMethod("New", func(info *types.PageInfo) *types.Redirect {
		info.Template = "Index"
		info.Assign("mod", "new")
		return nil
	})
}
