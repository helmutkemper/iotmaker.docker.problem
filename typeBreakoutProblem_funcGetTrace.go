package iotmaker_docker_problem

import (
	"log"
	"runtime"
)

func (e BreakoutProblem) getTrace() (file string, line int, funcName string, ok bool) {
	var pc uintptr

	pc, file, line, ok = runtime.Caller(1)
	if !ok {
		log.Printf("TraceToLog().error: runtime.Caller() unknown error")
		return
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		funcName = ""
		return
	}
	funcName = fn.Name()
	return
}
