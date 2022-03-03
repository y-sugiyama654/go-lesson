package main

import (
	"fmt"
	"regexp"
)

func main() {

	// 正規表現1
	// "apple"を判定したい場合
	// 先頭が"a"
	// 後尾が"e"
	// 先頭と後尾の間にa-zが1個以上存在する
	match, _ := regexp.MatchString("a([a-z]+)e", "applee")
	fmt.Println(match)

	// "a([a-z]+)e"の正規表現を何度も使いたい場合はMustCompileを使って変数に代入する
	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	// 正規表現2
	// "/view/test"を判定したい場合
	// ^先頭
	// "/"以降がeditかviewかsave
	// "/"以降がa-zA-Z0-9が1個以上存在する
	r2 := regexp.MustCompile("^/(edit|view|save|)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	// 正規表現にマッチしたらFindStringの引数内の文字列を表示する
	fmt.Println(fs)

	// 正規表現にマッチした際に、部分的にマッチした箇所を取り出したい場合
	fss := r2.FindStringSubmatch("/save/movie")
	fmt.Println(fss)
	fmt.Println(fss[0])
	fmt.Println(fss[1])
	fmt.Println(fss[2])
}
