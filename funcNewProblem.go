package iotmaker_docker_problem

func NewProblem(error, cause string) *BreakoutProblem {
	var problem = &BreakoutProblem{
		Err: error,
		Cau: cause,
	}

	return problem
}
