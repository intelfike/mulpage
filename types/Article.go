// 実行可能な構造体定義
// types.Packageに含まれる
package types

import (
	"errors"
)

// メソッドの定義
type Method func(PageInfo, *TplData) *Redirect

func (m Method) Exec(info PageInfo, tpl *TplData) (*Redirect, error) {
	if m == nil {
		fullm := info.FullMethod()
		return nil, errors.New(fullm + ":メソッドが正しく定義されていません")
	}
	return m(info, tpl), nil
}

// メソッドのラッパー
type Page struct {
	Index  int
	Key    string
	Name   string
	Method Method
}

func (a *Page) Exec(info PageInfo, tpl *TplData) (*Redirect, error) {
	return a.Method.Exec(info, tpl)
}
