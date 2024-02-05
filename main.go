package main

import (
	"proyecto-calculadora/BackEnd/api"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		api.IniciarAPI()
	}()

	wg.Wait()
}
