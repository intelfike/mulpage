package html

import (
	"bytes"
	"html/template"
	"io"
	"strconv"

	isear "github.com/intelfike/mulpage/contents/isear/page"
	"github.com/intelfike/mulpage/global"
	"github.com/intelfike/mulpage/proc/module/rand"
	"github.com/intelfike/mulpage/types"
)

func init() {
	// コンテンツのリストを定義
	var ContentList = map[string]types.ContentIfc{
		"isear": &isear.Content{},
	}
	global.Contents.Init("mulpage", ContentList)
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
	info.Assign("Rand", strconv.Itoa(rand.IntR()))

	// 関数実行
	redirect, err := global.Contents.Exec(info)
	if err != nil {
		return redirect, err
	}

	// テンプレート追加はPackage.Exec後
	info.AddTpl(info.LayoutPath(), info.PageTemplatePath())
	// テンプレート実行
	tpl, err := template.ParseFiles(info.TemplateFiles...)
	if err != nil {
		return redirect, err
	}
	err = tpl.Execute(w, info.AssignData)
	return redirect, err
}
