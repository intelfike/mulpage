package template

// import (
// 	tplproc "github.com/intelfike/mulpage/dev/proc/module/template"
// 	"github.com/intelfike/mulpage/dev/types"
// )

// func ParsePage(info *types.PageInfo) (string, error) {
// 	tpl := info.PageTemplatePath()
// 	s, err := tplproc.ParseFile(tpl, info)
// 	if err != nil {
// 		return "", err
// 	}
// 	info.PageBody = s

// 	layout := info.LayoutPath()
// 	s, err = tplproc.ParseFile(layout, info)
// 	if err != nil {
// 		return "", err
// 	}

// 	return s, nil
// }
