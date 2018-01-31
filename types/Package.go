package types

import (
	"fmt"

	"os"
)

// パッケージの定義
type Package struct {
	Name    string
	Before  Method
	After   Method
	Methods map[string]Method
}

type PackageIfc interface {
	Define(*Package)
}

func (pack *Package) Init(name string) {
	pack.Name = name
	pack.Methods = map[string]Method{}
}

func (ms *Package) SetMethod(key string, m Method) {
	if ms.Methods == nil {
		fmt.Println("types.Packageは必ずInit()してください\n例: pack.Init()")
		os.Exit(1)
	}
	if _, ok := ms.Methods[key]; ok {
		fmt.Println(key, ":メソッドは既に定義されています。")
	}
	ms.Methods[key] = m
}

func (pack *Package) Exec(tpl *TplData, info PageInfo) (*Redirect, error) {
	// 前置処理
	if red, _ := pack.Before.Exec(tpl, info); red != nil {
		return red, nil
	}
	// 実行
	method, ok := pack.Methods[info.Method]
	if !ok {
		return nil, nil
	}
	if red, err := method.Exec(tpl, info); red != nil {
		return red, err
	}
	// 後置処理
	red, _ := pack.After.Exec(tpl, info)
	return red, nil
}
