package main

import "fmt"
import "encoding/binary"
import "bytes"

func main() {

  N := 2147483647
  
  // if N not in range - throw exception
  if N < 1 && N > 2147483647 {
    fmt.Println ("Received N in invalid range")
  }
  
  number := uint32(N)
  var binBuf bytes.Buffer
  binLen := binary.PutUvarint(binBuf.Buffer(), number) // int
  
  yesBitPositionsArr := [33]int
  yesBitPositions := yesBitPositionsArr[:]
  var bitPositionsCell int
  
  for i:=0; i<(binLen-1); i+1 {
    
    if binBuf[i] == 1 {
      
      yesBitPositions = appendInt(yesBitPositions, i)
    
    }
    
    if cap(yesBitPositions) == len(yesBitPositions) {fmt.Println("Slice yesBitPositions is full!")}
    
    
  }
  
}
