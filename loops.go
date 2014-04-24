package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
)

func main() {
	t := time.Now()
	fmt.Println("seed",t.Unix())
	rand.Seed(int64(t.Unix()))
    for i:=1;i < 100;i++ {
    	j := rand.Intn(i)
    	f := float64(j)
    	fmt.Println("My favorite number is", i,j,math.Sqrt(f))
    }
}