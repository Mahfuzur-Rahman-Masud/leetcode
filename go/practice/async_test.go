package practice

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAsync(t *testing.T) {
	w := sync.WaitGroup{}
	w.Add(500)
	m := sync.Mutex{}
	sum := 0

	for x := 0; x < 500; x++ {
		go func() {
			m.Lock()
			sum++
			m.Unlock()
			w.Done()
		}()
	}

	w.Wait()
	fmt.Println(sum)
}

func BenchmarkAsync(t *testing.B) {
	w := sync.WaitGroup{}
	w.Add(500)
	m := sync.Mutex{}
	sum := 0

	for x := 0; x < 500; x++ {
		go func() {
			m.Lock()
			sum++
			m.Unlock()
			w.Done()
		}()
	}

	w.Wait()
	fmt.Println(sum)
}

func BenchmarkAsync2(t *testing.B) {
	w := sync.WaitGroup{}
	w.Add(500)
	sum := 0

	for x := 0; x < 500; x++ {
		go func() {
			sum++
			w.Done()
		}()
	}

	w.Wait()
	fmt.Println(sum)
}

func TestChanFanOut(t *testing.T) {
	a := make(chan string)
	b := make(chan string)
	c := make(chan string)
	ichan := make(chan string)
	fmt.Println("Cannels created!")

	go func() {
		for msg := range ichan {
			fmt.Println("New msg received")
			for i := 0; i < 3; i++ {
				switch i {
				case 0:
					a <- msg
				case 1:
					b <- msg
				case 2:
					c <- msg
				}
			}
		}
	}()

	fmt.Println("Creating subscriptions")

	go func() {
		for msg := range a {
			fmt.Println("Chan a:", msg)
		}
	}()

	go func() {
		for msg := range b {
			fmt.Println("Chan b:", msg)
		}
	}()

	go func() {

		for msg := range c {
			fmt.Println("Chan c:", msg)
		}
	}()

	fmt.Println("Writing to channel")

	ichan <- "Hello world"
	ichan <- "Hello universe"
	time.Sleep(time.Millisecond * 100)
}

func TestChanFanIn(t *testing.T) {
	pipe := make(chan string)
	i1 := make(chan string)
	i2 := make(chan string)
	i3 := make(chan string)

	go func() {
		for msg := range pipe {
			fmt.Println("In: ", msg)
		}
	}()

	go func() {
		for msg := range i1 {
			pipe <- "| i1 | " + msg
		}
	}()

	go func() {
		for msg := range i2 {
			pipe <- "| i2 | " + msg
		}
	}()

	go func() {
		for msg := range i3 {
			pipe <- "| i3 | " + msg
		}
	}()

	i1 <- "Hello world"
	i2 <- "Hi universe"
	i3 <- "Cow lives matter"
	i1 <- "World is created about 15b years ago"
	i2 <- "Universe is much wider than you think"
	i3 <- "Cows mow in the fild"
	i1 <- "Earth is the world we know as our home"
	close(i1)

	time.Sleep(time.Millisecond * 100)

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate some work
		fmt.Printf("Worker %d completed job %d\n", id, job)
		results <- job * 2
	}
}

func TestWorkerPool(t *testing.T) {
	const numJobs = 10
	const numWorkers = 3

	// Create job and result channels
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Create worker pool
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results)
		}(i)
	}

	// Send jobs to the workers
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results
	for result := range results {
		fmt.Printf("Result [%d]: %d\n", result/2, result)
	}
}
