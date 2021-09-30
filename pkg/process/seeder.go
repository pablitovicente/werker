package process

import (
	"math/rand"

	"github.com/pablitovicente/wrkpool/pkg/models"
)

func GenerateJob(ID int) models.Job {
	min := 10
	max := 48
	randPayload := uint(rand.Intn(max-min) + min)

	job := models.Job{
		ID:      ID,
		Payload: randPayload,
	}

	return job
}
