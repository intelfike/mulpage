package path

import "strings"

// パスをアプリケーションが利用可能な形にする
func ParseURIPath(path []string) []string {
	if path == nil || len(path) == 0 {
		// デフォルトの動作
		path = append(path, "_top")
		return path
	}
	if len(path) == 1 && strings.HasPrefix(path[0], "_") {
		// デフォルトの動作
		path = append(path, "Index")
	}
	return path
}

// パスをファイルパスで利用可能な形に変換する
func AvailePackageName(PackageName string) string {
	pn = strings.TrimPrefix(pn, "_")
	return pn
}
