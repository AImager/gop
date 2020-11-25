package gop

import "regexp"

var aspectMap = map[string][]AspectInterface{}

// get aspect
func getAspect(location string) []AspectInterface {
	return aspectMap[location]
}

// put match aspect
func putMatchAspect(location string) {
	aspectList := []AspectInterface{}
	for _, aspect := range aspectContainer {
		if isAspectMatch(aspect.GetAspectExpress(), location) {
			aspectList = append(aspectList, aspect)
		}
	}
	aspectMap[location] = aspectList
}

// util check if aspect match location
func isAspectMatch(aspectExpress, location string) bool {
	// regular
	pattern, err := regexp.Compile(aspectExpress)
	if err != nil {
		return false
	}
	return pattern.MatchString(location)
}
