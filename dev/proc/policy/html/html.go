// "github.com/intelfike/mulpage/dev/proc/html"
package html

import (
	"bytes"
	"html/template"
	"io"
	"strconv"

	isear "github.com/intelfike/mulpage/dev/contents/isear/page"
	"github.com/intelfike/mulpage/dev/global"
	"github.com/intelfike/mulpage/dev/ifc"
	"github.com/intelfike/mulpage/dev/proc/module/rand"
	"github.com/intelfike/mulpage/dev/types"
)

// コンテンツのリストを定義
var ContentList = map[string]ifc.Content{
	"isear": &isear.Content{},
}

func init() {
	for key, val := range ContentList {
		content := types.Content{}
		content.Init()
		global.Contents.SetContent(key, content)
		val.Init(content)
	}
}

// テンプレートをHTMLに変換して取得する
func Genelate(contents, packageName, method string) (string, *types.Redirect, error) {
	b := new(bytes.Buffer)
	redirect, err := Write(b, contents, packageName, method)
	return b.String(), redirect, err
}

// writerに変換済みテンプレートを書き出す
func Write(w io.Writer, contents, packageName, method string) (*types.Redirect, error) {
	// PageInfo 定義
	info := &types.PageInfo{}
	info.Init(contents, packageName, method)
	info.AddTpl(info.LayoutPath(), info.PageTemplatePath())
	info.Assign("Rand", strconv.Itoa(rand.IntR()))

	// 関数実行
	redirect, err := global.Contents.Exec(info)
	if err != nil {
		return redirect, err
	}
	// テンプレート実行
	tpl, err := template.ParseFiles(info.TemplateFiles...)
	if err != nil {
		return redirect, err
	}
	err = tpl.Execute(w, info.AssignData)
	return redirect, err
}
