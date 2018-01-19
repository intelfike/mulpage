package policy

// typesパッケージ内からはioパッケージを呼び出すべきでは無いため、PageInfo.Initではなくこちらに
// func CreatePageInfo(r *http.Request) (types.PageInfo, error) {
// 	info := types.PageInfo{"homepage", "top", "Index"} // デフォルト値
// 	// パスを取得
// 	path := httpio.ReadPath(r)
// 	if path != nil && !strings.HasPrefix(path[0], "_") {
// 		return info, errors.New("Method呼び出しではありません")
// 	}
// 	// ページ情報を定義
// 	info.Contents = httpio.ReadContents(r)
// 	if len(path) >= 1 {
// 		info.Package = strings.TrimPrefix(path[0], "_")
// 	}
// 	if len(path) >= 2 {
// 		info.Method = path[1]
// 	}
// 	return info, nil
// }
