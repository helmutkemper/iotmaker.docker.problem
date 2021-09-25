package iotmaker_docker_problem

import "fmt"

func ExampleBreakoutProblem_Trace() {
	var file string
	var line int
	var funcName string
	var ok bool

	var problem = NewProblem(
		"file not found",
		"file not found at path",
	)

	fmt.Printf("Error: %v\n", problem.Error())
	fmt.Printf("Cause: %v\n", problem.Cause())

	file, line, funcName, ok = problem.getTrace()

	fmt.Printf("File: %v\n", file)
	fmt.Printf("Line: %v\n", line)
	fmt.Printf("Func: %v\n", funcName)
	fmt.Printf("Ok: %v", ok)

	// Output:
	// Error: file not found
	// Cause: file not found at path
	// File: /Users/kemper/go/projetos/iotmaker.docker.problem/typeBreakoutProblem_funcTrace_test.go
	// Line: 19
	// Func: github.com/helmutkemper/iotmaker%2edocker%2eproblem.ExampleBreakoutProblem_Trace
	// Ok: true
}
