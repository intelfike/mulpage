package types

import (
	"errors"
	"fmt"
)

// パッケージの定義
type Package struct {
	Title   string
	Before  Method
	After   Method
	Methods map[string]Method
}

type PackageIfc interface {
	Define(*Package)
	Title() string
}

func (pack *Package) Init(title string) {
	pack.Title = title
	pack.Methods = map[string]Method{}
}

func (ms *Package) SetMethod(key string, m Method) {
	ms.Methods[key] = m
}

func (pack *Package) Exec(info *PageInfo) (*Redirect, error) {
	fmt.Println(pack)
	// 前置処理
	if red, _ := pack.Before.Exec(info); red != nil {
		return red, nil
	}
	// 実行
	method, ok := pack.Methods[info.Method]
	if !ok {
		return nil, errors.New(info.Method + ":メソッドは定義されていません")
	}
	if red, err := method.Exec(info); red != nil {
		return red, err
	}
	// 後置処理
	red, _ := pack.After.Exec(info)
	return red, nil
}
