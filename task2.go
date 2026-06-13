package main

import (
	"errors"
	"sort"
)

type PrintJob struct {
	ID        string
	Volume    float64
	Priority  int
	PrintTime int
}

type PrinterConstraints struct {
	MaxVolume float64
	MaxItems  int
}

type PrintingResult struct {
	PrintOrder []string
	TotalTime  int
}

func OptimizePrinting(printJobs []PrintJob, constraints PrinterConstraints) (PrintingResult, error) {
	if constraints.MaxVolume <= 0 {
		return PrintingResult{}, errors.New("max volume must be greater than 0")
	}
	if constraints.MaxItems <= 0 {
		return PrintingResult{}, errors.New("max items must be greater than 0")
	}

	jobs := make([]PrintJob, len(printJobs))
	copy(jobs, printJobs)

	for _, job := range jobs {
		if job.Volume <= 0 {
			return PrintingResult{}, errors.New("job volume must be greater than 0")
		}
		if job.PrintTime <= 0 {
			return PrintingResult{}, errors.New("job print time must be greater than 0")
		}
		if job.Priority < 1 || job.Priority > 3 {
			return PrintingResult{}, errors.New("job priority must be 1, 2 or 3")
		}
		if job.Volume > constraints.MaxVolume {
			return PrintingResult{}, errors.New("job volume exceeds printer max volume")
		}
	}

	sort.SliceStable(jobs, func(i, j int) bool {
		return jobs[i].Priority < jobs[j].Priority
	})

	var result PrintingResult
	var currentGroup []PrintJob
	currentVolume := 0.0

	for _, job := range jobs {
		canAddToGroup := len(currentGroup) < constraints.MaxItems &&
			currentVolume+job.Volume <= constraints.MaxVolume

		if !canAddToGroup {
			result.TotalTime += groupPrintTime(currentGroup)
			currentGroup = nil
			currentVolume = 0
		}

		currentGroup = append(currentGroup, job)
		currentVolume += job.Volume
		result.PrintOrder = append(result.PrintOrder, job.ID)
	}

	if len(currentGroup) > 0 {
		result.TotalTime += groupPrintTime(currentGroup)
	}

	return result, nil
}

func groupPrintTime(group []PrintJob) int {
	maxTime := 0
	for _, job := range group {
		if job.PrintTime > maxTime {
			maxTime = job.PrintTime
		}
	}
	return maxTime
}
