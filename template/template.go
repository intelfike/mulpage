package tpl

import (
	"html/template"
	"io"
)

// writerに変換済みテンプレートを書き出す
func Write(w io.Writer, data map[string]interface{}, filename ...string) error {
	tplfiles, err := template.ParseFiles(filename...)
	if err != nil {
		return err
	}
	err = tplfiles.Execute(w, data)
	return err
}
