package main

import (
	"fmt"
	"time"
  "math/rand"
)

func randomChromosome(length int) []int{
  rand.Seed(time.Now().Unix())
  numbers := make([]int, length)
  for i := 0; i < length; i++ {
    numbers[i] = rand.Intn(2) // Returns an integer in [0,2) = {0,1}
  }
  return numbers
}

func computeFitness(chromosome []int) int {
  ones := 0
  for i := 0; i < len(chromosome); i++ {
    if chromosome[i] == 1 {
      ones += 1
    }
  }
  return ones

}

func timeTrack(start time.Time, name string) {
  elapsed := time.Since(start).Seconds()
  fmt.Println(name,"took",elapsed,"seconds")
}


func main() {
  chromosome := randomChromosome(5)
  fmt.Println(chromosome,"->",computeFitness(chromosome))
}
