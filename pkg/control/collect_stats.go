package control

import "fmt"

func CollectStats(pendingJobs <-chan int) {
	for pending := range pendingJobs {
		fmt.Printf("CONTROL: Pening jobs: %d\n", pending)
	}
}
