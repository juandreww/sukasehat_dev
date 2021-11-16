package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  fmt.Printf("Spaceline \t\t Days \t Type \t\t Price \t Number of Days \t Departure Date \t\t Arrival Date \t\t\t Total Cost\n")
  fmt.Printf("=================================================================================================================================================================\n")
  for i := 1; i <= 10; i++ {
    var intSpaceline int = rand.Intn(3) + 1
    var daysSpent int = rand.Intn(20) + 20
    var singleOrRound int = rand.Intn(2) + 1

    spaceline := ""
    countTrip := ""

    switch intSpaceline {
    case 1:
      spaceline = "SpaceX\t\t"
    case 2:
      spaceline = "Virgin Galactic"
    case 3:
      spaceline = "Space Adventures"
    }

    switch singleOrRound {
    case 1:
      countTrip = "One-way"
    case 2:
      countTrip = "Round-trip"
    }

    var distanceToMars int = 62100000
    var baseSpeedSpaceline int = rand.Intn(15)
    var speedSpaceline int =  baseSpeedSpaceline  + 16
    var baseCostSpaceline int = 35 + baseSpeedSpaceline
    secondSpent := distanceToMars / speedSpaceline / 216000
    costSpaceline := baseCostSpaceline * singleOrRound
    departure := time.Date(2020, 10, 13, 0, 0, 0, 0, time.UTC)
    arrival := departure.AddDate(0, 0, secondSpent)

    fmt.Printf("%v \t %v \t %v \t $S %v \t %v \t\t\t %v \t %v \t $S %v\n", spaceline, daysSpent, countTrip, baseCostSpaceline , secondSpent, departure, arrival, costSpaceline)

  }
}