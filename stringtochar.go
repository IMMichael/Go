package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main() {
  
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Enter a string") 
  str, _ := reader.ReadString('\n')
   
  /* fmt.Println(str) */

  str = strings.Trim(str, "\n")

  size := len(str) - 1

  converted := strings.Split(str, "")
  for i:=0; i<size; i++ {
    fmt.Println(converted[i])
  }
  
}
