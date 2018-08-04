package main

import (
	"math/rand"
	"fmt"
)

type Person struct {
	isZombie bool
}

func New() *Person {
	person := &Person{}
	if rand.Intn(100) < 1 {
		person.isZombie = true
	}
	return person
}

func personDies(zombiePercent float64) bool {
	if rand.Intn(100) < int(zombiePercent) {
		return true
	} else {
		return false
	}
}

func becomesZombie(zombiePercent float64) bool {
	if rand.Intn(100) < int(zombiePercent + 4) / 2 {
		return true
	} else {
		return false
	}
}

func zombieDies(zombiePercent float64) bool {
	if rand.Intn(100) < int(zombiePercent) {
		return true
	} else {
		return false
	}
}

func getZombiePercent(population []Person) float64 {
	zombieCount := 0.0
	notZombieCount := 0.0
	for _, person := range population {
		if person.isZombie {
			zombieCount += 1
		} else {
			notZombieCount += 1
		}
	}
	return zombieCount / (zombieCount + notZombieCount)
}

func removePersonAtIndex(i int, a []Person) []Person {
	a[i] = a[len(a)-1]     // Copy last element to index i
	a[len(a)-1] = Person{} // Erase last element (write zero value)
	a = a[:len(a)-1]
	return a
}

func encounter(person1Index int, person2Index int, population []Person, zombiePercent float64) []Person {
	var zombie int
	var person int
	if population[person1Index].isZombie && !population[person2Index].isZombie {
		zombie = person1Index
		person = person2Index
	}
	if !population[person1Index].isZombie && population[person2Index].isZombie {
		zombie = person2Index
		person = person1Index
	}
	if zombie == 0 && person == 0 {
		return population
	}
	if becomesZombie(zombiePercent) {
		population[person].isZombie = true
	} else {
		if personDies(zombiePercent) {
			population = removePersonAtIndex(person, population)
		} else {
			if zombieDies(zombiePercent) {
				population = removePersonAtIndex(zombie, population)
			}
		}
	}
	return population
}

func main() {
	iterations := 0
	population := []Person{}
	for i := 0; i < 1000; i++ {
		population = append(population, *New())
	}
	percent := getZombiePercent(population)
	for i := 0; i < len(population); i++ {
		for j := 0; j < len(population); j++ {
			if i != j {
				iterations += 1
				population = encounter(i, j, population, percent)
				percent = getZombiePercent(population) * 100
				if percent == 0 || percent == 100 {
					goto End
				}
			}
			fmt.Printf("%.2f, %d\n", percent, len(population))
		}
	}
	End:
	fmt.Println(iterations)

}
