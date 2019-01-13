// 末端パッケージ
// http関連の入出力を共通化するためのパッケージ
package httpio

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// httpボディを読み取り、ボディのリーダーを閉じる
func ReadBody(r *http.Request) (string, error) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	return string(b), err
}

// 戻り値はパス
func ReadPath(r *http.Request) []string {
	path := r.URL.EscapedPath()
	// パスを取得
	path = strings.Trim(path, "/")
	if len(path) == 0 {
		return nil
	}
	return strings.Split(path, "/")
}

// コンテンツという名の、アプリケーションの動作を決める大まかなグループを取得
// ドメインの一番左側を取得する
func ReadContents(r *http.Request) string {
	host := strings.Split(r.Host, ".")
	return host[0]
}
// 引数の文字列を返す
func Write(w http.ResponseWriter, s string) error {
	_, err := w.Write([]byte(s))
	return err
}
// ファイルを読み取り、httpに書き込む
func WriteFile(w http.ResponseWriter, fileName string) error {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	conType := http.DetectContentType(b)
	switch {
	case strings.HasSuffix(fileName, ".css"):
		conType = "text/css"
	case strings.HasSuffix(fileName, ".js"):
		conType = "text/javascript"
	}
	w.Header().Set("Content-Type", conType)
	w.Write(b)
	return nil
}
