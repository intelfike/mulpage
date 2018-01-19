package page

import (
	"github.com/intelfike/mulpage/contents/isear/page/support"
	"github.com/intelfike/mulpage/contents/isear/page/top"
	"github.com/intelfike/mulpage/contents/isear/page/usage"
	"github.com/intelfike/mulpage/global"
	"github.com/intelfike/mulpage/types"
)

type Content struct{}

func (c *Content) Define(content *types.Content) {
	// パッケージのリストを定義
	var PackageList = map[string]types.PackageIfc{
		"top":     &top.Package{},
		"support": &support.Package{},
		"usage":   &usage.Package{},
	}
	content.Init("isear", PackageList)
	global.DefTplData.Assign("Packages", content.Packages)
}
