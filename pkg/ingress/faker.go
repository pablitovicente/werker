package ingress

import (
	"math/rand"

	"github.com/pablitovicente/wrkpool/pkg/models"
)

func SeedSampleData(numJobs int, jobs chan<- models.Job) {
	for j := 1; j <= numJobs; j++ {
		newJob := generateJob(j)
		jobs <- newJob
	}
	close(jobs)
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
