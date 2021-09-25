package iotmaker_docker_problem

import (
	"log"
	"runtime"
)

type BreakoutProblem struct {
	Err string
	Cau string

	file     string
	line     int
	funcName string
	ok       bool
}

func (e BreakoutProblem) Trace() (file string, line int, funcName string, ok bool) {
	return e.file, e.line, e.funcName, e.ok
}

func (e BreakoutProblem) Error() string {
	return e.Err
}

func (e BreakoutProblem) getTrace() (file string, line int, funcName string, ok bool) {
	var pc uintptr

	pc, file, line, ok = runtime.Caller(2)
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

func (e BreakoutProblem) Cause() string {
	return e.Cau
}
