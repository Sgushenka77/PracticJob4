package main
import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func worker(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- j * j
	}
	wg.Done()
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	
	for work := 1; work <= 3; work++ {
		wg.Add(1)
		go worker(jobs, results)
	}
	
	for i := 10; i <= 19; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	close(results)
	
	for res := range results {
		fmt.Println(res)
	}
}