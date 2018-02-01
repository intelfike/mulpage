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
		if v, ok := con.Children[info.Package].Articles[info.Method]; ok {
			tpl.Assign("Title", v.Name)
		}
		tpl.Assign("Info", info)
		return nil
	}
}
