package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// resp, _ := http.Get("https://umayadia-apisample.azurewebsites.net/api/persons/Shakespeare")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	// URLのパース
	base, _ := url.Parse("http://example.com")
	// パラメータのパース
	reference, _ := url.Parse("/test?a=1&b=2")
	// URLにパラメータをセット
	endpoint := base.ResolveReference(reference).String()

	// リクエストの作成
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-Match", `Q/xyz`)

	// クエリの確認
	q := req.URL.Query()

	// パラメータの追加
	q.Add("c", "3&%")

	fmt.Println(q)
	fmt.Println(q.Encode())

	req.URL.RawQuery = q.Encode()

	// クライアントの作成
	var client *http.Client = &http.Client{}
	// クライアントの実行
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
