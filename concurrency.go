package main

import(
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

func Announce(message string, delay time.Duration) {
    // go func() {
        // time.Sleep(delay)
        fmt.Println("yo")
    // }()  // Note the parentheses - must call the function.
}	


func main() {
	for i:=0;i < 100;i++ {
		Announce("its "+strconv.Itoa(i)+" oclock", time.Duration(rand.Intn(i+1)) )
	}
}