package main

import(
  "fmt"
  "strings"

)

func main() {

  s := "A very very very long sentence."
  
  // TODO: Length of string
  fmt.Println(len(s))

  // Iterate over each character
  for _, ch := range s {
    fmt.Print(string(ch), ",")
  }

  // TODO: Use math operators  == != <>
  fmt.Print("\ndog is: ")
  fmt.Println("dog" == "Dog")

  fmt.Print("Cat is: ")
  fmt.Println("Cat" == "Cat")

  fmt.Print("Greater than is: ")
  fmt.Println("a" < "b")

  fmt.Print("Smaller than is: ")
  fmt.Println("a" > "b")
  
  //TODO: Comparing two strings

  result := strings.Compare("cat", "dog") 
  fmt.Println(result)

  result = strings.Compare("cat", "cat") 
  fmt.Println(result)

  // TODO: EqualFold tests using unicode case folding
  fmt.Println(strings.EqualFold("Hello", "hello"))

  // TODO: ToUpper ToLower
  s1 := strings.ToUpper(s)
  s2 := strings.ToLower(s)
  s3 := strings.(s)
  fmt.Println(s1)
  fmt.Println(s2)

}
