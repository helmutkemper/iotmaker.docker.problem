package iotmaker_docker_problem

type Problem interface {
	Error() string
	Cause() string
	Trace() (file string, line int, funcName string, ok bool)
}
