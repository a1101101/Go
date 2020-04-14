/*
An interface is a group of methods which are to be attached to whatever data type you want to use.
Declaration of an interface requires only the name of its method and the type of the returned value.
Instantiation of interface I means an instance of type I having the methods of I is being made.
You can assign an instance of any type to an instance of an interface.

Below, U is instantiated as "u:=...".
Then, u U is assigned to interface i I as "i=u".

Note that the T() method is called as a method for i I in func main, but defined as a method for u U.
*/
package main

import (
	"fmt"
    "reflect"
)

type U int

type I interface {
	T()
}

func main() {	
    //Instansiation of U
    u := U(3)
    fmt.Println(reflect.TypeOf(u))
    
    //Instantiation of I
    var i I
    
    //Assign u to i
	i = u
	
    i.T()
}

//Define T() as a method for U
func (u U) T(){
    fmt.Println(reflect.TypeOf(u))
}
