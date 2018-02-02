package html

import (
	"bytes"
	"html/template"
	"io"
	"strconv"

	"github.com/intelfike/module/rand"
	isear "github.com/intelfike/mulpage/contents/isear/page"
	www "github.com/intelfike/mulpage/contents/www/page"
	"github.com/intelfike/mulpage/global"
	"github.com/intelfike/mulpage/types"
)

func init() {
	// コンテンツリストを定義
	global.App.Init("app", "mulpage")

	wwwCon := global.App.NewChild("www")
	wwwCon.FallDown(&www.Content{})

	isearCon := global.App.NewChild("isear")
	isearCon.FallDown(&isear.Content{})
}

// テンプレートをHTMLに変換して取得する
// いらない
func Genelate(info types.PageInfo) (string, *types.Redirect, error) {
	b := new(bytes.Buffer)
	redirect, err := Write(b, info)
	return b.String(), redirect, err
}

// writerに変換済みテンプレートを書き出す
func Write(w io.Writer, info types.PageInfo) (*types.Redirect, error) {
	tpl := &types.TplData{}
	tpl.Init()
	tpl.SetPageInfo(info)
	tpl.Assign("Rand", strconv.Itoa(rand.IntR()))

	// 関数実行
	redirect, err := global.App.Exec(tpl, info)
	if redirect != nil || err != nil {
		return redirect, err
	}

	// テンプレート追加はPackage.Exec後
	tpl.AddTpl(tpl.LayoutPath(info), tpl.PageTemplatePath(info))
	// テンプレート実行
	tplfiles, err := template.ParseFiles(tpl.TemplateFiles...)
	if err != nil {
		return redirect, err
	}
	err = tplfiles.Execute(w, tpl.AssignData)
	return redirect, err
}
