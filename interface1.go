/*
An interface is a group of methods which are to be attached to a certain data type.
Declaring an interface requires only the name of its method and the type of the returned value.
Instantiation of I creates an instance of type I with methods of I.

Below, U is instantiated as "u:=...".
Then, u U is assigned to interface i I as "i=u".

Note that the T() method is called as a method for i I in func main, but defined as a method for u U.
This is also mentioned in the official tutorial "Tour of Go" as an implicit implementation of an interface. 
*/
package main

import (
	"fmt"
	"reflect"
)

//Declaring type U
type U int

//Declaring interface I
type I interface {
    T()
}

func main() {	
    //Instantiation of U
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
