package process

import (
	"time"

	"github.com/pablitovicente/wrkpool/pkg/models"
)

// The actual worker; multiple instances of this will be created
func Worker(workerId int, jobs <-chan models.Job, results chan<- models.Job) {
	for j := range jobs {
		start := time.Now()
		j.Result = j.Executor.Exec(j.Payload)

		j.ProcStats = append(j.ProcStats, models.WorkerStats{
			ID:       workerId,
			ExecTime: time.Since(start),
		})

		results <- j
	}
}
