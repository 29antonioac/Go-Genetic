package main

import (
	"fmt"
	"time"
  "math/rand"
)

// Create a random chromosome
func randomChromosome(length int) []int{
  numbers := make([]int, length)
  for i := 0; i < length; i++ {
    numbers[i] = rand.Intn(2) // Returns an integer in [0,2) = {0,1}
  }
  return numbers
}

// Computes maxOnes fitness
func computeFitness(chromosome []int) int {
  ones := 0
  for i := 0; i < len(chromosome); i++ {
    if chromosome[i] == 1 {
      ones += 1
    }
  }
  return ones
}

// Mutate a chromosome
func mutate(chromosome []int) []int {
  mutationPoint := rand.Intn(len(chromosome))

  result := make([]int, len(chromosome))
  copy(result,chromosome)

  result[mutationPoint] = (result[mutationPoint] + 1) % 2 // Change result[mutationPoint] from 0 to 1 and viceversa

  return result
}

// Mutate all chromosomes in a population
func mutatePopulation(population [][]int) [][]int {
  result := make([][]int, len(population))

  for i := range result {
    result[i] = mutate(population[i])
  }
  return result
}

// Track the time elapsed by a function (use with defer)
func timeTrack(start time.Time, name string) {
  elapsed := time.Since(start).Seconds()
  fmt.Println(name,"took",elapsed,"seconds")
}


func main() {
  rand.Seed(time.Now().Unix())
  chromosome := randomChromosome(5)
  chr_mutated := mutate(chromosome)
  fmt.Println(chromosome,"->",computeFitness(chromosome),"->",chr_mutated)



  population := make([][]int, 6)
  for i := range population {
    population[i] = randomChromosome(6)
  }
  fmt.Println(population,"->",mutatePopulation(population))



}
