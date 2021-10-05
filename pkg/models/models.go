package models

import "time"

type Job struct {
	ID        int
	Payload   uint
	Result    interface{}
	ProcStats []WorkerStats
	Executor  Step
}

type Step interface {
	Exec(uint) uint
}

type WorkerStats struct {
	ID          int
	ExecTime    time.Duration
	OperationID int
}

type Stats struct {
	OperationID int
	JobID       int
	ExecTime    time.Duration
}

var WorkerOperations = map[int]string{
	1: "Calculate Fibonacci",
	2: "TBD",
}
