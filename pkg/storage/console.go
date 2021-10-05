package storage

import (
	"fmt"

	"github.com/pablitovicente/wrkpool/pkg/models"
)

func Log(numJobs int, pendingJobs chan<- int, results <-chan models.Job) {
	completed := 0
	for a := 1; a <= numJobs; a++ {
		res := <-results
		completed++
		// Update stats
		pendingJobs <- numJobs - completed
		fmt.Printf("STORAGE: Job ID: %d fib(%d) === %d (Processing Stats: %+v)\n", res.ID, res.Payload, res.Result, res.ProcStats)
		fmt.Println("................................................")
		fmt.Println("")
	}
}
