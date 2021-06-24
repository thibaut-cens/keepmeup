package main

import (
  "time"
  "fmt"
  "math/rand"
  "github.com/go-vgo/robotgo"
)

func wake(r *rand.Rand) {
    fmt.Println("Pressing numlock")
    robotgo.KeyToggle("numpad_lock")

    // Generate a random number between 
    var kdelay = r.Intn(400) + 100
    fmt.Printf("Waiting between keystrokes for %d ms\n", kdelay)
    time.Sleep(time.Duration(kdelay) * time.Millisecond)

    fmt.Println("Pressing numlock")
    robotgo.KeyToggle("numpad_lock")
}

func main() {
  var r = rand.New(rand.NewSource(time.Now().UnixNano()))
  for {
    var ldelay = r.Intn(1000) + 9000
    fmt.Printf("Waiting between loop iterations for %d ms\n", ldelay)
    time.Sleep(time.Duration(ldelay) * time.Millisecond)
    
    wake(r)
  }
}
