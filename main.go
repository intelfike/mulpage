// UNIX哲学に出来る限り従う方向で！
// isearとホームページを作りたい
// lazyの構成を参考に
package main

import (
	"flag"
	"fmt"
	"net/http"

	isear "github.com/intelfike/mulpage/contents/isear/page"
	"github.com/intelfike/mulpage/global"
	httpio "github.com/intelfike/mulpage/io/http"
	"github.com/intelfike/mulpage/policy"
	htmlproc "github.com/intelfike/mulpage/policy/html"
	"github.com/intelfike/mulpage/types"
)

var port = flag.String("http", ":80", "HTTP port number.")

func init() { // コンテンツのリストを定義
	flag.Parse()

	// コンテンツリストを定義
	var ContentList = map[string]types.ContentIfc{
		"isear": &isear.Content{},
	}
	// グローバル変数の設定
	global.App.Init("mulpage", ContentList)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		info, err := policy.CreatePageInfo(r)
		if err != nil {
			httpio.WriteFile(w, "public"+r.URL.Path)
			return
		}
		// HTMLを生成して送信
		redirect, err := htmlproc.Write(w, info)
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
