package main

import (
	"math/rand"
	"sync"
	"time"
)

func main() {
	pokemonList := []string{"Pikachu", "Charmander", "Squirtle", "Bulbasaur", "Jigglypuff"}
	cond := sync.NewCond(&sync.Mutex{})
	pokemon := ""

	// Consumer
	go func() {
		cond.L.Lock()
		defer cond.L.Unlock()

		// waits until Pikachu appears
		for pokemon != "Pikachu" {
			println("1 Saw " + pokemon)
			cond.Wait()
		}
		println("1 Caught " + pokemon)
	}()

	// Consumer
	go func() {
		cond.L.Lock()
		defer cond.L.Unlock()

		// waits until Pikachu appears
		for pokemon != "Pikachu" {
			println("2 Saw " + pokemon)
			cond.Wait()
		}
		println("2 Caught " + pokemon)
	}()

	// Producer
	go func() {
		// Every 1ms, a random Pokémon appears
		for range 100 {
			time.Sleep(time.Millisecond)

			cond.L.Lock()
			pokemon = pokemonList[rand.Intn(len(pokemonList))]
			cond.L.Unlock()

			cond.Signal()
		}
	}()

	time.Sleep(100 * time.Millisecond) // lazy wait
}
