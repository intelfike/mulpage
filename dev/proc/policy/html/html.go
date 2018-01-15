// "github.com/intelfike/mulpage/dev/proc/html"
package html

import (
	"bytes"
	"html/template"
	"io"
	"strconv"

	_ "github.com/intelfike/mulpage/dev/contents/isear/page"
	"github.com/intelfike/mulpage/dev/global"
	"github.com/intelfike/mulpage/dev/proc/module/rand"
	"github.com/intelfike/mulpage/dev/types"
)

// テンプレートをHTMLに変換して取得する
func Genelate(contents, packageName, method string) (string, *types.Redirect, error) {
	b := new(bytes.Buffer)
	redirect, err := Write(b, contents, packageName, method)
	return b.String(), redirect, err
}

func Write(w io.Writer, contents, packageName, method string) (*types.Redirect, error) {
	info := types.DefaultPageInfo(contents, packageName, method)
	info.Assign("Rand", strconv.Itoa(rand.IntR()))
	global.PageInfo = info
	redirect, err := global.Contents.Exec(info)
	if err != nil {
		return redirect, err
	}
	tpl, err := template.ParseFiles(info.LayoutPath(), info.PageTemplatePath())
	if err != nil {
		return redirect, err
	}
	err = tpl.Execute(w, info.AssignData)
	return redirect, err
}
