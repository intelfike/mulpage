// リクエスト情報から取り出したデータ
// リクエストによって定義される
// ページ関連のファイルを取得するための情報
package types

import (
	"errors"
	"net/http"
	"strings"

	"github.com/intelfike/ideapouch/policy/jsondb"
	httpio "github.com/intelfike/mulpage/io/http"
)

type PageInfo struct {
	Contents string // dev/contents/*
	Group    string // dev/contents/*/page/X
	Page     string // dev/contents/*/page/X/X.go.Function

	Method   string
	Data     map[string]string
	PostData map[string]string

	JsonDB *jsondb.JsonDB
}

func NewPageInfo(r *http.Request) (PageInfo, error) {
	info := PageInfo{}
	err := info.Init(r)
	return info, err
}

// 初期化処理
func (info *PageInfo) Init(r *http.Request) error {
	// コンテンツを取得
	last_domain := httpio.ReadContents(r)
	if strings.Contains(last_domain, ":") {
		last_domain = strings.Split(last_domain, ":")[0]
	}
	info.Contents = last_domain
	// パスを取得
	path := httpio.ReadPath(r)
	if path != nil && !strings.HasPrefix(path[0], "_") {
		return errors.New("Method呼び出しではありません")
	}
	// ページ情報を定義
	if len(path) >= 1 {
		info.Group = strings.TrimPrefix(path[0], "_")
	}
	if len(path) >= 2 {
		info.Page = path[1]
	}

	// データを取得
	info.Method = r.Method
	r.ParseForm()
	info.Data = map[string]string{}

	for key, _ := range r.Form {
		info.Data[key] = r.Form.Get(key)
	}
	if r.Method == "POST" {
		info.PostData = map[string]string{}
		for key, _ := range r.PostForm {
			info.PostData[key] = r.PostForm.Get(key)
		}
	}

	// jsondbを定義
	info.JsonDB = jsondb.NewJsonDB("contents/ideapouch/data")
	info.JsonDB.SetIndent("\t")

	return nil
}

// ピリオド区切りでメソッド名のフルパスを表示
func (pi *PageInfo) FullMethod() string {
	return strings.Join([]string{pi.Contents, pi.Group, pi.Page}, ".")
}

// 相対パスを取得する
// func (pi *PageInfo) ContentsPath() string {
// 	path := filepath.Join("contents", pi.Contents)
// 	return path
// }
// func (pi *PageInfo) PackagePath() string {
// 	path := filepath.Join(pi.ContentsPath(), "page", pi.Group)
// 	return path
// }
