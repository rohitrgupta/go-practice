package main

func Multiple3And5(num int) int {
  sum := 0
  for i := 0; i < num ; i++ {
    if i % 3 == 0  || i % 5  == 0 {
      sum += i
    }
  }

  return sum
}
