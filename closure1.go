/*
Go's closure is similiar to C++'s lambda expression such like "auto f = [&]([args]){...}".
*/
package main

import "fmt"

func main() {
    sum := 0
    //f references sum outside its body 
    var f = func(x int) int{
        sum += x
        return sum
    }
    
    for i := 0; i < 10; i++ {
	fmt.Println(f(i))
        fmt.Println(sum)
    }
}
