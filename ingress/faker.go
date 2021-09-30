package ingress

import (
	"github.com/pablitovicente/wrkpool/models"
	"github.com/pablitovicente/wrkpool/processing"
)

func SeedSampleData(numJobs int, jobs chan<- models.Job) {
	for j := 1; j <= numJobs; j++ {
		newJob := processing.GenerateJob(j)
		jobs <- newJob
	}
	close(jobs)
}
