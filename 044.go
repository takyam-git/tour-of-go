package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  var f func(num int) int
  memo := map[int]int{}
  n := 0
  f = func(num int) int {
    if num < 2 {
        return num
    }
    //check memo
    if _, exists := memo[num]; !exists {
      // calculate fibonacci number if did not calculate
      memo[num] = f(num - 2) + f(num - 1)
    }
    
    return memo[num]
  }
  return func() int {
    result := f(n)
    n++
    return result
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
      fmt.Println(f())
  }
}
