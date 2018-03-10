// メソッドによって定義されるテンプレート関連情報
// テンプレートを構築するための情報はここに。
package types

import "path/filepath"

type TplData struct {
	// デフォルトでMethod名と同じ、変更可能
	Template string
	Layout   string

	// テンプレートファイルの位置
	RootPath     string
	TemplatePath string // ルートからの相対パス
	LayoutPath   string // ルートからの相対パス

	// テンプレートファイルをリクエスト時に追加するため
	TemplateFiles []string
	// 主にテンプレートに渡すデータたち
	AssignData map[string]interface{}
}

func (tpl *TplData) Init() {
	*tpl = TplData{
		Template:      "",
		Layout:        "default",
		TemplateFiles: []string{},
		AssignData:    map[string]interface{}{},
	}
}

// init以降のイニシャライザ
func (tpl *TplData) SetPageInfo(info PageInfo) {
	tpl.Template = info.Article
	tpl.Assign("Title", info.Package+" | "+info.Contents)
}

// データを追加
func (tpl *TplData) AddTpl(fileNames ...string) {
	tpl.TemplateFiles = append(tpl.TemplateFiles, fileNames...)
}
func (tpl *TplData) Assign(key string, value interface{}) {
	tpl.AssignData[key] = value
}

// 相対パスを取得する
func (tpl *TplData) GetTemplatePath() string {
	path := filepath.Join(tpl.RootPath, tpl.TemplatePath, tpl.Template+".tpl")
	return path
}

func (tpl *TplData) GetLayoutPath() string {
	path := filepath.Join(tpl.RootPath, tpl.LayoutPath, tpl.Layout+".tpl")
	return path
}
