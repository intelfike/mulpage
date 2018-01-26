// UNIX哲学に出来る限り従う方向で！
// isearとホームページを作りたい
// lazyの構成を参考に
package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/intelfike/mulpage/policy"
)

var port = flag.String("http", ":80", "HTTP port number.")

func init() { // コンテンツのリストを定義
	flag.Parse()

	http.HandleFunc("/", policy.Listener)
}

func main() {
	fmt.Printf("Start HTTP Server localhost%s\n", *port)
	fmt.Println(http.ListenAndServe(*port, nil))
}
