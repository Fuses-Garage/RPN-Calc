package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int//式の要素数
	var err error//エラーを入れる変数
	fmt.Print("入力される演算子と被演算子の数の合計を入力してください->")
	_, err = fmt.Scan(&n)
	for err != nil {//エラーが起きたらもう一度
		fmt.Println("整数を入力してください")
		fmt.Print("入力される演算子と被演算子の数の合計を入力してください->")
		_, err = fmt.Scan(&n)
	}
	stack := make([]float64, 0, n)
	var siki string
	fmt.Print("逆ポーランド表記で式を入力してください(カンマ区切り）->")
	fmt.Scanln(&siki)
	sikiarray := strings.Split(siki, ",")
	for len(sikiarray) != n {//要素数が違ったらもう一度
		fmt.Println(siki)
		fmt.Printf("要素数が異なります 期待：%d、実際:%d\n", n, len(sikiarray))
		fmt.Print("逆ポーランド表記で式を入力してください(カンマ区切り）->")
		_, err = fmt.Scanln(&siki)
		sikiarray = strings.Split(siki, ",")
	}
	for i := 0; i < n; i++ {

		var val float64
		val, err = strconv.ParseFloat(sikiarray[i], 64)//要素を少数に変換しようとする
		if err == nil {//要素が値なら
			stack = append(stack, val)//スライスに積む
		} else {//値でない（演算子）なら
			if len(stack) >= 2 {//値が前に2つ以上あれば
				a := stack[len(stack)-2]
				b := stack[len(stack)-1]
				ans := 0.0
				switch sikiarray[i] {//文字列でスイッチ
				case "+":
					ans = a + b
				case "-":
					ans = a - b
				case "*":
					ans = a * b
				case "/":
					ans = a / b
				default://四則演算子以外が入っていたら
					panic("不正な演算子が入力されました")//強制終了
				}
				stack = stack[:len(stack)-2]//スライスを2つ削る
				stack = append(stack, ans)//演算結果を積む
			} else {
				panic("式が不正です")//強制終了
			}

		}
	}
	if len(stack)==1{//スタックの中身が1つだけなら
		fmt.Printf("%s=%f\n", siki, stack[0])//結果出力
	}else{//被演算子が2つ以上残ってたら
		panic("式が不正です")//強制終了
	}
	
}
