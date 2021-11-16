package main

import (
  "fmt"
  "math/rand"
)

var era = "AD"

func main() {
  for i := 1; i <= 10; i++ {
    year := rand.Intn(20) + 2010
    month := rand.Intn(12) + 1
    daysInMonth := 31
    switch month {
    case 2:
      if year%4 == 0 {
        daysInMonth = 29
      } else {
        daysInMonth = 28
      }
    case 4, 6, 9, 11:
      daysInMonth = 30
    }
    day := rand.Intn(daysInMonth) + 1
    fmt.Println(era, year, month, day)
  }
}