// UNIX哲学に出来る限り従う方向で！
// isearとホームページを作りたい
// lazyの構成を参考に
package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	httpio "github.com/intelfike/mulpage/io/http"
	htmlproc "github.com/intelfike/mulpage/proc/policy/html"
	pathpol "github.com/intelfike/mulpage/proc/policy/path"
)

var port = flag.String("http", ":80", "HTTP port number.")

func init() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// パスを取得
		path := httpio.ReadPath(r)
		// パスをアプリケーション向けに変換
		path = pathpol.ParseURIPath(path)
		if len(path) == 0 || !strings.HasPrefix(path[0], "_") {
			err := httpio.WriteFile(w, "public"+r.URL.Path)
			if err == nil {
				// ファイルが見つからない場合の処理
				return
			}
			return
		}
		if len(path) == 1 {
			http.Redirect(w, r, path[0], 307)
			return
		}
		// HTMLを生成して送信
		contents := httpio.ReadContents(r)
		PackageName := pathpol.AvailePackageName(path[0])
		redirect, err := htmlproc.Write(w, contents, PackageName, path[1])
		if err != nil {
			result := "エラー:" + fmt.Sprint(err) + "\n"
			result += "エラーメッセージを記録して管理者に報告してください。"
			httpio.Write(w, result)
			return
		}
		// 必要ならリダイレクトする
		redirect.Exec(w, r)
	})
}

func main() {
	fmt.Printf("Start HTTP Server localhost%s\n", *port)
	fmt.Println(http.ListenAndServe(*port, nil))
}