package main

import (
    "fmt"
)

type sayer interface {
    say(a string) string
}

type AA struct{
    name string
}


func (this *AA) say(a string) string {
    out := "==========>AA " + a
    fmt.Println(out)
    return out
}


type BB struct{
    *AA
    age int
}

func (this *BB) say(a string) string {
    out := "==========>BB " + a
    fmt.Println(out)
    return out
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
    o1   := obj1.say("monkey")
    obj2 := ObjectFactory(0)
    o2   := obj2.say("business")
    fmt.Println("outputs o1,o2",o1,o2)
}