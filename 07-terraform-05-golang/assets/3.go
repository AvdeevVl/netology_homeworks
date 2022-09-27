package main

import "fmt"

func main() {

 var s []int
 
 for i := 1; i <= 100; i++ {
  if i % 3 == 0 {
   s = append(s, i)
  }  
 }
 fmt.Println(s)
}