// メソッドを階層的に格納するための構造体。
// 構造体内の変数を直接弄くった場合は、動作を保証できません。
package types

import (
	"errors"
	"fmt"

	"os"
)

// パッケージの定義
type Package struct {
	Index    int
	Key      string
	Before   Method
	After    Method
	Children map[string]Executable
}

type Executable interface {
	Exec(PageInfo, *TplData) (*Redirect, error)
}

func NewPackage() *Package {
	pack := new(Package)
	pack.Init()
	return pack
}

// 構造体を初期化する
func (pack *Package) Init() {
	pack.Children = map[string]Executable{}
}

type Definer interface {
	Define(*Package)
}

// 子パッケージの定義を実行する
func (pack *Package) FallDown(pis ...Definer) {
	for _, pi := range pis {
		pi.Define(pack)
	}
}

// 新しい子パッケージを作成し、追加する
// SetChildのラッパ
func (pack *Package) NewChild(key string) *Package {
	if _, ok := pack.Children[key]; ok {
		fmt.Println(key, ":重複したキーによる子パッケージは登録できません。\n違う名前で登録してください。")
		return nil
	}
	child := NewPackage()
	child.Index = len(pack.Children)
	child.Key = key

	pack.SetChild(key, child)
	return child
}

// 子パッケージを追加する
func (pack *Package) SetChild(key string, child *Package) {
	pack.setExecutable(key, child)
}

// 実行可能なメソッドを追加する
func (pack *Package) SetMethod(key string, m Method) {
	atc := &Page{len(pack.Children), key, m}
	pack.setExecutable(key, atc)
}

// 子を追加する
func (pack *Package) setExecutable(key string, child Executable) {
	if pack.Children == nil {
		fmt.Println("types.Packageは必ずInit()してください\n例: pack.Init(\"web\")")
		os.Exit(1)
	}
	if _, ok := pack.Children[key]; ok {
		fmt.Println(key, ":重複したキーによるメソッドは登録できません。\n違う名前で登録してください。")
	}

	pack.Children[key] = child
}

// リクエストに対してプログラムを実行する
// Beforeでリダイレクト→即座にリダイレクトするかどうか
// 階層が2の場合、Before1 => Before2 => Exec2 => After2 => After1 という順番で実行する
func (pack *Package) Exec(info PageInfo, tpl *TplData) (*Redirect, error) {
	return pack.exec(info, tpl, "/root", 0)
}
func (pack *Package) exec(info PageInfo, tpl *TplData, path string, depth int) (*Redirect, error) {
	if depth >= 3 {
		return nil, errors.New("URLのパスが長すぎます")
	}
	// 前置処理
	if red, _ := pack.Before.Exec(info, tpl); red != nil {
		return red, nil
	}
	key := ""
	switch depth {
	case 0:
		key = info.Contents
	case 1:
		key = info.Group
	case 2:
		key = info.Page
	}
	// 実行
	atc, ok := pack.Children[key]
	if !ok {
		return nil, errors.New("エラー:" + path + "において、" + key + "は定義されていません")
	} else {
		fmt.Println(key, atc)
	}
	// 単純な実行
	switch t := atc.(type) {
	case *Package:
		red, err := t.exec(info, tpl, path+"/"+key, depth+1)
		if red != nil || err != nil {
			return red, err
		}
	default:
		red, err := atc.Exec(info, tpl)
		if red != nil || err != nil {
			return red, err
		}
	}

	// 後置処理
	red, _ := pack.After.Exec(info, tpl)
	return red, nil
}

// キーを追加した順番でソートし、配列で返す
func (pack *Package) ChildrenArray() []Executable {
	arr := make([]Executable, len(pack.Children))
	for key, val := range pack.Children {
		switch v := val.(type) {
		case *Page:
			v.Key = key
			arr[v.Index] = v
		case *Package:
			v.Key = key
			arr[v.Index] = v
		}
	}
	return arr
}
