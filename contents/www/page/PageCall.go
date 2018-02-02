package page

import (
	"github.com/intelfike/mulpage/contents/www/page/top"
	"github.com/intelfike/mulpage/types"
)

type Content struct{}

func (c *Content) Define(con *types.Package) {
	con.Init("content", "homepage")

	topPack := con.NewChild("top")
	topPack.FallDown(&top.Package{})

	con.Before = func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.Assign("Packages", con.ChildrenArray())
		tpl.Assign("Package", con.Children[info.Package])
		if v, ok := con.Children[info.Package]; ok {
			if v1, ok1 := v.Articles[info.Article]; ok1 {
				tpl.Assign("Title", v1.Name)
			}
		}
		tpl.Assign("Info", info)
		return nil
	}

	con.After = func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.Assign("Template", tpl.Template)
		return nil
	}
}
