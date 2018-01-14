// "github.com/intelfike/mulpage/dev/proc/html"
package html

import (
	"bytes"
	"html/template"
	"io"

	"github.com/intelfike/mulpage/dev/contents/isear/page"
	"github.com/intelfike/mulpage/dev/global"
	"github.com/intelfike/mulpage/dev/types"
)

func init() {
	page.Call()
}

// テンプレートをHTMLに変換して取得する
func Genelate(contents, pageName, method string) (string, *types.Redirect, error) {
	b := new(bytes.Buffer)
	redirect, err := Write(b, contents, pageName, method)
	return b.String(), redirect, err
}

func Write(w io.Writer, contents, pageName, method string) (*types.Redirect, error) {
	info := types.DefaultPageInfo(contents, pageName, method)
	global.PageInfo = info
	redirect, err := global.Contents.Exec(info)
	if err != nil {
		return redirect, err
	}
	tpl, err := template.ParseFiles(info.LayoutPath(), info.PageTemplatePath())
	if err != nil {
		return redirect, err
	}
	err = tpl.Execute(w, info)
	return redirect, err
}