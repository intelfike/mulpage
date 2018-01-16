package page

import (
	"github.com/intelfike/mulpage/dev/contents/isear/page/support"
	"github.com/intelfike/mulpage/dev/contents/isear/page/top"
	"github.com/intelfike/mulpage/dev/contents/isear/page/usage"
	"github.com/intelfike/mulpage/dev/ifc"
	"github.com/intelfike/mulpage/dev/types"
)

// パッケージのリストを定義
var PackageList = map[string]ifc.Package{
	"top":     &top.Package{},
	"support": &support.Package{},
	"usage":   &usage.Package{},
}

type Content struct{}

var _ ifc.Content = &Content{}

func (c *Content) Init(content types.Content) {
	for key, val := range PackageList {
		p := types.Package{}
		p.Init()
		content.SetPackage(key, p)
		val.Init(p)
	}
}
