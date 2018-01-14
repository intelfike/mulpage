package top

import (
	"github.com/intelfike/mulpage/dev/global"
	"github.com/intelfike/mulpage/dev/types"
)

type Page struct{}

func (p *Page) Init(page types.Page) {
	page.SetMethod("Index", func() *types.Redirect {
		info := global.PageInfo
		info.Title = "isearweb"
		info.Assign("mod", "modmod")
		return nil
	})

	page.SetMethod("New", func() *types.Redirect {
		info := global.PageInfo
		info.Template = "Index"
		info.Assign("mod", "new")
		return nil
	})
}
