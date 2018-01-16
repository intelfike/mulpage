package types

import (
	"errors"
)

// パッケージの定義
type Package struct {
	Before  Method
	After   Method
	Methods map[string]Method
}

func (pack *Package) Init() {
	pack.Methods = map[string]Method{}
}
func (pack *Package) ExecBefore(info *PageInfo) (*Redirect, error) {
	if pack.Before == nil {
		return nil, nil
	}
	return pack.Before.Exec(info)
}

func (pack *Package) ExecAfter(info *PageInfo) (*Redirect, error) {
	if pack.After == nil {
		return nil, nil
	}
	return pack.After.Exec(info)
}

func (pack *Package) Exec(info *PageInfo) (*Redirect, error) {
	// 前置処理
	if red, err := pack.ExecBefore(info); red != nil || err != nil {
		return red, err
	}
	// 実行
	method, ok := pack.Methods[info.Method]
	if !ok {
		return nil, errors.New(info.Method + ":メソッドは定義されていません")
	}
	if red, err := method.Exec(info); red != nil || err != nil {
		return red, err
	}
	// 後置処理
	red, err := pack.ExecBefore(info)
	return red, err
}

func (ms *Package) SetMethod(key string, m Method) {
	ms.Methods[key] = m
}
