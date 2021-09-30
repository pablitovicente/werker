package main

import (
	"fmt"

	"github.com/pablitovicente/wrkpool/pkg/control"
	"github.com/pablitovicente/wrkpool/pkg/ingress"
	"github.com/pablitovicente/wrkpool/pkg/models"
	"github.com/pablitovicente/wrkpool/pkg/process"
)

func main() {
	// Configuration for the number of fake jobs
	const numJobs = 32
	// Number of workers to process the jobs
	const numOfWorkers = 32

	// All required channels
	jobs := make(chan models.Job, numJobs)
	results := make(chan models.Job, numJobs)
	pendingJobs := make(chan int)

	// Create a new worker pool
	process.NewPool(numOfWorkers, jobs, results)
	// Generate fake jobs
	go ingress.SeedSampleData(numJobs, jobs)
	// Stats collection
	go control.CollectStats(pendingJobs)

	// Just to show non-blocking nature of this code
	// go func() {
	// 	for {
	// 		fmt.Println("I just run every 500ms to show the non blocking nature of this code")
	// 		time.Sleep(500 * time.Millisecond)
	// 	}
	// }()

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
