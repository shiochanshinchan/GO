package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards :=rand.Intn(100)
  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.") 
    fmt.Println("eludedGuards has a value of", eludedGuards)  
  } else {
    fmt.Println("Plan a better disguise next time?")
    fmt.Println("eludedGuards has a value of", eludedGuards) 
  }
  fmt.Println("isHeistOn is currently:",isHeistOn)

  if eludedGuards >= 70 && isHeistOn == true {
    fmt.Println("Grab and GO!")
  }else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")  
    fmt.Println("What's the combo to this lock again??") 
  }
  leftSafely := rand.Intn(5)
  if isHeistOn {
    switch leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("Looks like you tripped an alarm... run?")
      case 1: 
        isHeistOn = false 
        fmt.Print("Turns out vault doors don't open from the inside...") 
      case 2: 
        isHeistOn = false 
        fmt.Print("Turns out vault doors don't open from the inside...") 
      case 3: 
        isHeistOn = false 
        fmt.Print("Turns out vault doors don't open from the inside...")
      default:
        fmt.Println("Start the getaway car!")   
        amtStolen := 10000 + rand.Intn(1000000)  
        fmt.Println("$", amtStolen, "not bad!") 
    }
  }

}
