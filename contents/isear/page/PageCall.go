package page

import (
	"github.com/intelfike/mulpage/contents/isear/page/install"
	"github.com/intelfike/mulpage/contents/isear/page/support"
	"github.com/intelfike/mulpage/contents/isear/page/top"
	"github.com/intelfike/mulpage/contents/isear/page/usage"
	"github.com/intelfike/mulpage/types"
)

type Content struct{}

func (c *Content) Define(con *types.Package) {
	con.Init("content", "isear")

	topPack := con.NewChild("top")
	topPack.FallDown(&top.Package{})

	usagePack := con.NewChild("usage")
	usagePack.FallDown(&usage.Package{})

	supportPack := con.NewChild("support")
	supportPack.FallDown(&support.Package{})

	installPack := con.NewChild("install")
	installPack.FallDown(&install.Package{})

	con.Before = func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.Assign("Packages", con.ChildrenArray())
		tpl.Assign("Package", con.Children[info.Package])
		tpl.Assign("Info", info)
		return nil
	}

	// type PA struct {
	// 	Package string
	// 	Ins     types.PackageIfc
	// 	Text    string
	// } // Package Assign

	// PArr := []PA{
	// 	PA{"top", &top.Package{}, ""},
	// 	PA{"usage", &usage.Package{}, ""},
	// 	PA{"support", &support.Package{}, ""},
	// 	PA{"install", &install.Package{}, ""},
	// } // Package Array
	// PList := map[string]types.PackageIfc{}

	// for _, v := range PArr {
	// 	PList[v.Package] = v.Ins
	// }

	// // パッケージのリストを定義
	// content.Init("isear", PList)

	// for n, v := range PArr {
	// 	name := content.Packages[v.Package].Name
	// 	PArr[n].Text = name
	// }

	// content.Before = func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
	// 	tpl.Assign("Packages", PArr)
	// 	tpl.Assign("Package", content.Packages[info.Package])
	// 	tpl.Assign("Info", info)
	// 	return nil
	// }
}
