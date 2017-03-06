package main

import "fmt"

func main() {
  for i:=0 ; i<2 ; i++ {
    var name string
    fmt.Scanf("%s", &name)
    fmt.Println("Hello!", name)
  }
  fmt.Println("Now, say number!")
  
  var number int
  fmt.Scanf("%i", &number)
  switch number {
    case 0: fmt.Println("Zero")
    case 1: fmt.Println("One")
    case 2: fmt.Println("Two")
    case 3: fmt.Println("Three")
    case 4: fmt.Println("Four")
    case 5: fmt.Println("Five")
    case 6: fmt.Println("Six")
    case 7: fmt.Println("Seven")
    case 8: fmt.Println("Eight")
    case 9: fmt.Println("Nine")
    default: fmt.Println("What I've got is not a number")
  }
}
