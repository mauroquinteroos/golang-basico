package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func execute_concurrency() {
	start := time.Now()
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	fmt.Printf("Tiempo inicial: %v\n", start)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readBook(i, wg, m)
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("Tiempo total: %v\n", duration)
}

func readBook(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
	FinishBook(id, m)
	delay := rand.Intn(800)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	wg.Done()
}
