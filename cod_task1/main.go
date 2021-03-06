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
  //binLen := binary.PutUvarint(binBuf.Write([]byte()), uint64(number)) // int
  
  //var binArr [32]byte
  
  err := binary.Write(&binBuf, binary.LittleEndian, N)
  if err != nil {
    fmt.Println("binary.Write failed:", err)
  }
  
  binLen := len(binBuf)
  
  var yesBitPositionsArr [33]int
  yesBitPositions := yesBitPositionsArr[:]
  var bitPositionsCell int
  
  for i:=0; i<(binLen-1); i++ {
    
    if binBuf[i] == 1 {
      
      yesBitPositions = appendInt(yesBitPositions, i)
    
    }
    
    if cap(yesBitPositions) == len(yesBitPositions) {fmt.Println("Slice yesBitPositions is full!")}
    
    
  }
  
}
