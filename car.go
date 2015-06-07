// source: http://mattn.kaoriya.net/software/lang/go/20121213175242.htm
// ascii art: http://yaruo.b4t.jp/index.php?id=HukuTemp/%E3%81%9D%E3%81%AE%E4%BB%96%E6%B1%8E%E7%94%A8AA/%E4%B9%97%E3%82%8A%E7%89%A9%E3%83%BB%E3%83%A1%E3%82%AB/%E8%87%AA%E5%8B%95%E8%BB%8A.mlt
package main

import (
	"time"
	"fmt"
	"strings"
)

func main() {
	p := [][][]string {
		{
			{
				" 　　　　　　　　　　　　　　　　　　　　____＿,,,、、､､､､､､､､､､､､､,,,,,,,,,,,,,＿_＿_",
				" 　　　　　　　　　　　　　　　　　／ﾞ´￣￣￣￣７　 , -ｰｰｰｰ‐ｰ,┌ーーーー､｀､",
				" 　　　　　　　　　　　　　　 　／　　　　　 　　／／　　　　　　 　l l　　　　 　　 l ﾞ,",
				" 　　　　　　　　　　　　　　／　　　　　　 　／／__　　　　　 　 　 l l　　　　 　　│ﾞ,",
				" 　　　　　　　　 　 _ , ､-´ｰｰｰｰｰｰｰｰ→´／ﾆ___）_________________l│__________,,,､」　l7",
				" 　　　　 _,､-‐･ﾞ´,=ﾆニ＞　　　,　- ｰ ﾞ ´　ｌ　　　　　 　 　 　 ___ l　　　　　　 　　　 │",
				" 　 　 /-,_________￣￣,´-､,･´　　　　　　 　ｌ　　　　　　　 　　｀ｰ´l 　　　　　　　　 　〔l",
				" 　　/（ノ　　,_______￣l,_,,ノ　 　,　----､　ﾛ.ｌ　　　　　　　　　　　　l　　 　　　/,ニニ､ l",
				" 　　ト　‐====,ｰｰ┘------,/･´_ﾆ_ﾞヽヽ,. ｌ　　　　　　　　　　　　l　　　　　ｌ / Vヾ,ヽｌ",
				" 　　ll二ll_____│￣l　l￣l　 /　 ,´ ∨｀, ﾞ, l│　　　　　　　　　　　ﾉ　　　　　ｌ l＞o＜l ﾘ",
				" 　 └-ｰ====ｰﾆﾆ､二二 _l　 l,＞()＜l ｌ ｌﾞｰｰｰｰｰｰ-_-__二二二二ニニ/　ﾞ､.∧ ,ﾉ/",
				" 　 　　　･､｀ー´ノ　￣￣　 ｉ　　､ ∧丿ﾉ ￣￣￣￣　　　　　　　　　　　　ヽ､,,｀二ノ",
				" 　　　　　　｀ﾞﾞﾞﾞ　　　　　 　 ｀･､　二､･´",
			},
		},
	}

	x := 100
	for _, ll := range p {
		for _, l := range ll {
			for n := range l {
				l[n] = strings.Repeat(" ", x) + l[n]
			}
		}
	}
	x = 0

	for {
		fmt.Print("\x0c")
		for _, pp := range p {
			for _, l := range pp[x % len(pp)] {
				if x < len(l) {
					fmt.Println(string(l[x:]))
				} else {
					fmt.Println()
				}

			}
		}
		time.Sleep(time.Second / 50)
		if x++; x > 240 {
			break
		}
	}
}

