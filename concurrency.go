package main

// import(
// 	"fmt"
// 	"time"
// 	"math/rand"
// 	"strconv"
// )

// func Announce(message string, delay time.Duration) {
//     // go func() {
//         // time.Sleep(delay)
//         fmt.Println("yo")
//     // }()  // Note the parentheses - must call the function.
// }	


// func main() {
// 	for i:=0;i < 100;i++ {
// 		Announce("its "+strconv.Itoa(i)+" oclock", time.Duration(rand.Intn(i+1)) )
// 	}
//  time.Sleep(time.Duration(12 * 1e9))
// }

import (
    "fmt"
    "time"
)

func simulateEvent(name string, timeInSecs int64) { 
    // sleep for a while to simulate time consumed by event
    fmt.Println("Started ", name, ": Should take", timeInSecs, "seconds.")
    t := time.Now().Unix();
    time.Sleep(time.Duration(timeInSecs * 1e9))
    fmt.Println("Finished ", name, "in",time.Now().Unix()-t, "seconds")
}

func main() {
    go simulateEvent("100m sprint", 10) //start 100m sprint, it should take 10 seconds
    go simulateEvent("Long jump", 6) //start long jump, it should take 6 seconds
    go simulateEvent("High jump", 3) //start high jump, it should take 3 seconds
    time.Sleep(time.Duration(12 * 1e9))
}