// メソッドによって定義されるテンプレート関連情報
// テンプレートを構築するための情報はここに。
package types

import (
	"io"
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"

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
	assignData map[string]interface{}

	debugPrintText string
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
		assignData:    map[string]interface{}{},
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
	tpl.assignData[key] = value
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
	return mtpl.Write(w, tpl.assignData, tpl.TemplateFiles...)
}

func (tpl *TplData) DebugPrint(s string) {
	tpl.debugPrintText += s
}

func (tpl *TplData) GetDebugPrintText() string {
	return tpl.debugPrintText
}

// form.yamlを読み取る
// data = formから受け取ったkey-valueのデータ
func (tpl *TplData) LoadFormYaml(data map[string]interface{}, yamlfile string, inheritKeys ...string) error {
	// ファイルを読み取る
	// Decoderを使わないのは、あれだから。
	b, err := ioutil.ReadFile(filepath.Join(tpl.pagePath, yamlfile))
	if err != nil {
		return err
	}
	// yamlをデコードする
	form := make(map[string]interface{})
	err = yaml.Unmarshal(b, &form)
	if err != nil {
		return err
	}

	// データを整形する
	items := form["items"].(map[interface{}]interface{})
	for name, i_item := range items {
		item := i_item.(map[interface{}]interface{})
		d, ok := data[name.(string)]
		if !ok {
			d = item["default"]
		}
		item["value"] = d
		// 引き継ぐキーをアサイン
		for _, key := range inheritKeys {
			item[key] = tpl.assignData[key]
		}
		// 代入する
		items[name] = item
	}

	// TODO: バリデーション機能を作る

	tpl.Assign("FormData", items)

	return nil
}
