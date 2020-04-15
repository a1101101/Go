/*
Simple closure program
A closure is a func object which references variables outside its body.
It is an equivalent of "auto f = [&]([args]){...}" in C++.
*/
package main

import "fmt"

func main() {
    sum := 0
    //closure f which refers to sum outside its body 
    var f = func(x int) int{
        sum += x
        return sum
    }
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
        fmt.Println(sum)
	}
}