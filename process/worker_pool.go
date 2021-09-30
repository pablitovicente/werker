package process

import "github.com/pablitovicente/wrkpool/models"

func NewPool(numOfWorkers int, jobs, results chan models.Job) {
	for w := 1; w <= numOfWorkers; w++ {
		go Worker(w, jobs, results)
	}
}
