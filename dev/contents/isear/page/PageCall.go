package page

import (
	"github.com/intelfike/mulpage/dev/contents/isear/page/top"
	"github.com/intelfike/mulpage/dev/global"
	"github.com/intelfike/mulpage/dev/types"
)

// types.PageBaseを埋め込んだ./*/Pageを宣言してください。
var Pages = map[string]interface {
	Init(types.Page)
}{
	"top": &top.Page{},
}

func init() {
	ps := types.Pages{}
	global.Contents.SetPages("isear", ps)
	for key, val := range Pages {
		p := types.Page{}
		ps.SetPage(key, p)
		val.Init(p)
	}
}
