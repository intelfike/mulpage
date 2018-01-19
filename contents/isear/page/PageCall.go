package page

import (
	"github.com/intelfike/mulpage/contents/isear/page/install"
	"github.com/intelfike/mulpage/contents/isear/page/support"
	"github.com/intelfike/mulpage/contents/isear/page/top"
	"github.com/intelfike/mulpage/contents/isear/page/usage"
	"github.com/intelfike/mulpage/types"
)

type Content struct{}

func (c *Content) Define(content *types.Content) {
	// パッケージのリストを定義
	var PackageList = map[string]types.PackageIfc{
		"top":     &top.Package{},
		"support": &support.Package{},
		"usage":   &usage.Package{},
		"install": &install.Package{},
	}
	content.Init("isear", PackageList)

	// content.Before = func(tpl *types.TplData, info types.PageInfo) *types.Redirect {

	// 	return nil
	// }
}
