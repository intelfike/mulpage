// メソッドによって定義されるテンプレート関連情報
package types

import "path/filepath"

type TplData struct {
	// デフォルトでMethod名と同じ、変更可能
	Template string // dev/contents/*/page/X/template/*.tpl
	Layout   string

	// テンプレートファイルを追加するため
	TemplateFiles []string
	// 主にテンプレートに渡すデータたち
	AssignData map[string]interface{}
}

func (rd *TplData) Init() {
	*rd = TplData{
		Template:      "",
		Layout:        "default",
		TemplateFiles: []string{},
		AssignData:    map[string]interface{}{},
	}
}

// init以降のイニシャライザ
func (rd *TplData) SetPageInfo(info PageInfo) {
	rd.Template = info.Article
	rd.Assign("Title", info.Package+" | "+info.Contents)
}

// データを追加
func (rd *TplData) AddTpl(fileNames ...string) {
	rd.TemplateFiles = append(rd.TemplateFiles, fileNames...)
}
func (rd *TplData) Assign(key string, value interface{}) {
	rd.AssignData[key] = value
}

// 相対パスを取得する
func (rd *TplData) PageTemplatePath(info PageInfo) string {
	path := filepath.Join(info.PackagePath(), "template", rd.Template+".tpl")
	return path
}

func (rd *TplData) LayoutPath(info PageInfo) string {
	path := filepath.Join(info.ContentsPath(), "template", "layout", rd.Layout+".tpl")
	return path
}
