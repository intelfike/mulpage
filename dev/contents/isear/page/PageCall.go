package page

import (
	"github.com/intelfike/mulpage/dev/contents/isear/page/top"
	"github.com/intelfike/mulpage/dev/global"
	"github.com/intelfike/mulpage/dev/types"
)

// ./*/Pageをここに書いてください
var Packages = map[string]interface {
	Init(types.Package)
}{
	"top": &top.Package{},
}

func init() {
	ps := types.Content{}
	for key, val := range Packages {
		p := types.Package{}
		ps.SetPackage(key, p)
		val.Init(p)
	}
	global.Contents.SetContent("isear", ps)
}
