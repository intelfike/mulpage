// メソッドによって定義されるテンプレート関連情報
// テンプレートを構築するための情報はここに。
package types

import (
	"io"
	"path/filepath"

	mtpl "github.com/intelfike/mulpage/template"
)

type TplData struct {
	// デフォルトでMethod名と同じ、変更可能
	layout string
	page   string
	parts  []string

	// テンプレートファイルの位置
	layoutPath string
	pagePath   string
	partsPath  string

	// テンプレートファイルをリクエスト時に追加するため
	TemplateFiles []string
	// 主にテンプレートに渡すデータたち
	AssignData map[string]interface{}
}

func NewTplData() *TplData {
	tpl := new(TplData)
	tpl.Init()
	return tpl
}

func (tpl *TplData) Init() {
	*tpl = TplData{
		layout:        "default",
		page:          "",
		parts:         []string{},
		TemplateFiles: []string{},
		AssignData:    map[string]interface{}{},
	}
}

func (tpl *TplData) SetPage(page string) {
	tpl.page = page
}
func (tpl *TplData) SetLayout(layout string) {
	tpl.layout = layout
}
func (tpl *TplData) AddParts(fileNames ...string) {
	for _, v := range fileNames {
		name := tpl.partsPath + v
		tpl.parts = append(tpl.parts, name)
	}
}

func (tpl *TplData) Assign(key string, value interface{}) {
	tpl.AssignData[key] = value
}

// テンプレートファイルを追加 プロジェクトルートからのパス
func (tpl *TplData) AddFiles(fileNames ...string) {
	tpl.TemplateFiles = append(tpl.TemplateFiles, fileNames...)
}

func (tpl *TplData) SetLayoutPath(path string) {
	tpl.layoutPath = path
}
func (tpl *TplData) SetPagePath(path string) {
	tpl.pagePath = path
}
func (tpl *TplData) SetPartsPath(path string) {
	tpl.partsPath = path
}

// writerにテンプレートを書き出す
func (tpl *TplData) Write(w io.Writer) error {
	layout := filepath.Join(tpl.layoutPath, tpl.layout)
	page := filepath.Join(tpl.pagePath, tpl.page)

	tpl.AddFiles(layout, page)
	tpl.AddFiles(tpl.parts...)
	return mtpl.Write(w, tpl.AssignData, tpl.TemplateFiles...)
}
