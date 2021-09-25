package iotmaker_docker_problem

type BreakoutProblem struct {
	Err string
	Cau string

	file     string
	line     int
	funcName string
	ok       bool
}
