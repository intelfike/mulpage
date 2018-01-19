package policy

import (
	"errors"
	"net/http"
	"strings"

	httpio "github.com/intelfike/mulpage/io/http"
	"github.com/intelfike/mulpage/types"
)

// Method呼び出しでなければnil
func CreatePageInfo(r *http.Request) (types.PageInfo, error) {
	info := types.PageInfo{"homepage", "top", "Index"} // デフォルト値
	// パスを取得
	path := httpio.ReadPath(r)
	if path != nil && !strings.HasPrefix(path[0], "_") {
		return info, errors.New("Method呼び出しではありません")
	}
	// ページ情報を定義
	info.Contents = httpio.ReadContents(r)
	if len(path) >= 1 {
		info.Package = strings.TrimPrefix(path[0], "_")
	}
	if len(path) >= 2 {
		info.Method = path[1]
	}
	return info, nil
}
