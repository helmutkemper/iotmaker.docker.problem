package iotmaker_docker_problem

func (e BreakoutProblem) Trace() (file string, line int, funcName string, ok bool) {
	return e.file, e.line, e.funcName, e.ok
}
