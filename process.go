package gop

// process before call
func beforeProcessed(point *JoinPoint, location string) bool {
	for _, aspect := range aspectMap[location] {
		if !aspect.Before(point) {
			return false
		}
	}
	return true
}

// process after call
func afterProcessed(point *JoinPoint, location string) {
	for i := len(aspectMap[location]) - 1; i >= 0; i-- {
		aspect := aspectMap[location][i]
		aspect.After(point)
	}
}

// process after return, use defer
func finallyProcessed(point *JoinPoint, location string) {
	for i := len(aspectMap[location]) - 1; i >= 0; i-- {
		aspect := aspectMap[location][i]
		aspect.Finally(point)
	}
}
