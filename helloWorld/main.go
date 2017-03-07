package main

import "fmt"

func main() {
  
  // Loop section
  for i:=0 ; i<2 ; i++ {
    fmt.Print("String",i + 1,":")
    var name string
    fmt.Scanf("%s", &name)
    fmt.Println("Hello!", name)
  }
  
  // Switch section
  fmt.Println("Now, send me a number!")
  var number int
  fmt.Scanf("%d", &number)
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
  
  // Arrays/Slices
  month := [...]string {1:"January", 2:"February", 3:"March", 4:"April", 5:"May", 6:"June", 7:"July", 8:"August", 9:"September", 10:"October", 11:"November", 12:"December"}
  Q1 := month[1:4]
  Q2 := month[4:7]
  Q3 := month[7:10]
  Q4 := month[10:13]
  winter := month[1:3]
  spring := month[3:6]
  summer := month[6:9]
  autumn := month[9:13]
  
  fmt.Println(Q1)
  fmt.Println(Q2)
  fmt.Println(Q3)
  fmt.Println(Q4)
  
  fmt.Println(winter)
  fmt.Println(spring)
  fmt.Println(summer)
  fmt.Println(autumn)
  
}
