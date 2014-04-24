package main

import (
    "fmt"
)

type sayer interface {
    say()
}

type AA struct{
    name string
}

/*
func (this *AA) say(){
    fmt.Println("==========>AA")
}
*/

type BB struct{
    *AA
    age int
}
func (this *BB) say(){
    fmt.Println("==========>BB")
}

func ObjectFactory(typeNum int) sayer {
    if typeNum ==1 {
        return new(AA)
    }else{
        return new(BB)
    }
}

func main() {
    obj1 := ObjectFactory(1)
    obj1.say()
    obj2 := ObjectFactory(0)
    obj2.say()
}