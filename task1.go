package main
import (
	"fmt";
	"sync";
	"time"
)
var wg sync.WaitGroup
func main() {
	go countToFive()
	wg.Add(1)
	wg.Wait()
	fmt.Println("Подсчёт завершён")
}

func countToFive() {
	for i := 1; i <= 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i)
	}
	wg.Done()
}