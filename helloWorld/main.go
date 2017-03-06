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
    case 1: fmt.Println("Zero")
    case 2: fmt.Println("Zero")
    case 3: fmt.Println("Zero")
    case 4: fmt.Println("Zero")
    case 5: fmt.Println("Zero")
    case 6: fmt.Println("Zero")
    case 7: fmt.Println("Zero")
    case 8: fmt.Println("Zero")
    case 9: fmt.Println("Zero")
  }
}
