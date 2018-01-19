package types

import (
	"errors"
)

// メソッドの定義
type Method func(*TplData, PageInfo) *Redirect

func (m Method) Exec(tpl *TplData, info PageInfo) (*Redirect, error) {
	if m == nil {
		fullm := info.FullMethod()
		return nil, errors.New(fullm + ":メソッドが正しく定義されていません")
	}
	return m(tpl, info), nil
}
