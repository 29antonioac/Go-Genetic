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

// Crossover two chromosomes
func crossover(chromosome, another []int) [][]int{
  crossover_point := rand.Intn(len(chromosome) - 2)

  interval := 1 + rand.Intn(len(chromosome) - crossover_point - 1)

  fmt.Println(crossover_point, interval)

  new_chromosome  := chromosome[:crossover_point]
  new_another     := another[:crossover_point]

  new_chromosome  = append(new_chromosome,another[crossover_point:crossover_point + interval + 1]...)
  new_another     = append(new_another,chromosome[crossover_point:crossover_point + interval + 1]...)

  new_chromosome  = append(new_chromosome,chromosome[crossover_point + interval + 1:]...)
  new_another     = append(new_another,another[crossover_point + interval + 1:]...)

  return [][]int {new_chromosome, new_another}


}

func timeMutations(iterations int, individual []int) {
  defer timeTrack(time.Now(), "Mutations")

  for i := 0; i < iterations; i++ {
    mutate(individual)
  }
  return

}

func main() {
  rand.Seed(time.Now().Unix())
  // chromosome := randomChromosome(5)
  // chr_mutated := mutate(chromosome)
  // fmt.Println(chromosome,"->",computeFitness(chromosome),"->",chr_mutated)
  //
  //
  //
  // population := make([][]int, 6)
  // for i := range population {
  //   population[i] = randomChromosome(6)
  // }
  // fmt.Println(population,"->",mutatePopulation(population))
  //
  // fmt.Println(population[:2],"->",crossover(population[0],population[1]))

  length := 16
  iterations := 100000
  top_length := 32768

  for length < top_length {
    individual := randomChromosome(length)
    fmt.Print("Go-BitString ",length,", ")
    timeMutations(iterations,individual)
    length *= 2
  }





}
