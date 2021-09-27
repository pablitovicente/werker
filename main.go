package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pablitovicente/wrkpool/models"
)

// The actual worker; multiple instances of this will be created
func worker(workerId int, jobs <-chan models.Job, results chan<- models.Job) {
	for j := range jobs {
		start := time.Now()
		j.Result = fib(j.Payload)

		j.ProcStats = append(j.ProcStats, models.WorkerStats{
			ID:       workerId,
			ExecTime: time.Since(start),
		})

		results <- j
	}
}

func main() {
	// Configuration for the number of fake jobs
	const numJobs = 32
	// Number of workers to process the jobs
	const numOfWorkers = 32

	// All required channels
	jobs := make(chan models.Job, numJobs)
	results := make(chan models.Job, numJobs)
	pendingJobs := make(chan int)

	// Spawn all the configured workers
	for w := 1; w <= numOfWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Generate fake jobs
	for j := 1; j <= numJobs; j++ {
		newJob := generateJob(j)
		jobs <- newJob
	}
	// All jobs sent no need to keep the channel open
	close(jobs)

	// Just a listener of job completion to show some stats
	go func(pendingJobs <-chan int) {
		for pending := range pendingJobs {
			fmt.Printf("Pening jobs: %d\n", pending)
		}
	}(pendingJobs)

	// Just to show non-blocking nature of this code
	go func() {
		for {
			fmt.Println("I just run every 500ms to show the non blocking nature of this code")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Print results of jobs
	proccessed := 0
	for a := 1; a <= numJobs; a++ {
		res := <-results
		proccessed++
		// Update stats
		pendingJobs <- numJobs - proccessed
		fmt.Printf("Job ID: %d fib(%d) === %d (Processing Stats: %+v)\n", res.ID, res.Payload, res.Result, res.ProcStats)
		fmt.Println("................................................")
		fmt.Println("")
	}
}

func generateJob(ID int) models.Job {
	min := 10
	max := 48
	randPayload := uint(rand.Intn(max-min) + min)

	job := models.Job{
		ID:      ID,
		Payload: randPayload,
	}

	return job
}

// The infamously slow recursive Fibonacci algo
func fib(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
