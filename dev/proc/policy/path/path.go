package path

import "strings"

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
func AvailePageName(pageName string) string {
	pageName = strings.TrimPrefix(pageName, "_")
	pageName = strings.ToLower(pageName)
	return pageName
}
