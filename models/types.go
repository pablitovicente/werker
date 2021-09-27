package models

import "time"

type Job struct {
	ID        int
	Payload   uint
	Result    interface{}
	ProcStats []WorkerStats
}

type WorkerStats struct {
	ID          int
	ExecTime    time.Duration
	OperationID int
}

// For general stats
type Stats struct {
	OperationID int
	JobID       int
	ExecTime    time.Duration
}

var WorkerOperations = map[int]string{
	1: "Calculate Fibonacci",
	2: "TBD",
}
