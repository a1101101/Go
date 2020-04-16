/*
関数オブジェクト(クロージャー)を返す関数adder()
クロージャーは自身のスコープ外(環境)の変数を参照できる
以下ではadder()内でクロージャー"func(x int) int"が"sum"を参照している

Go言語では、関数オブジェクトを引数にとる、返り値にとる、変数に代入する、のいずれも可能。
*/
package main

import "fmt"

//adder() returns a closure
func adder() func(int) int {
	sum := 0
    //closure references sum
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i),neg(-2*i),)
	}
}
