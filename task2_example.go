package main

import "fmt"

func RunTask2Examples() {
	test1Jobs := []PrintJob{
		{ID: "M1", Volume: 100, Priority: 1, PrintTime: 120},
		{ID: "M2", Volume: 150, Priority: 1, PrintTime: 90},
		{ID: "M3", Volume: 120, Priority: 1, PrintTime: 150},
	}

	test2Jobs := []PrintJob{
		{ID: "M1", Volume: 100, Priority: 2, PrintTime: 120},
		{ID: "M2", Volume: 150, Priority: 1, PrintTime: 90},
		{ID: "M3", Volume: 120, Priority: 3, PrintTime: 150},
	}

	test3Jobs := []PrintJob{
		{ID: "M1", Volume: 250, Priority: 1, PrintTime: 180},
		{ID: "M2", Volume: 200, Priority: 1, PrintTime: 150},
		{ID: "M3", Volume: 180, Priority: 2, PrintTime: 120},
	}

	constraints := PrinterConstraints{
		MaxVolume: 300,
		MaxItems:  2,
	}

	fmt.Println("Task 2:")
	printOptimizationResult("Test 1 (same priority):", test1Jobs, constraints)
	printOptimizationResult("Test 2 (different priorities):", test2Jobs, constraints)
	printOptimizationResult("Test 3 (constraint limits):", test3Jobs, constraints)
}

func printOptimizationResult(title string, jobs []PrintJob, constraints PrinterConstraints) {
	result, err := OptimizePrinting(jobs, constraints)
	if err != nil {
		fmt.Println(title)
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(title)
	fmt.Println("Print order:", result.PrintOrder)
	fmt.Println("Total time:", result.TotalTime, "minutes")
	fmt.Println()
}
