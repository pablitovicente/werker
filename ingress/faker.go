package ingress

import (
	"github.com/pablitovicente/wrkpool/models"
	"github.com/pablitovicente/wrkpool/process"
)

func SeedSampleData(numJobs int, jobs chan<- models.Job) {
	for j := 1; j <= numJobs; j++ {
		newJob := process.GenerateJob(j)
		jobs <- newJob
	}
	close(jobs)
}
