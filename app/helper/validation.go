package helper

import "regexp"

func ValidationAirlineCode(airlineCode string) bool {
	// NOTE: Standard International Carrier Code
	coderegex := regexp.MustCompile("^[a-zA-Z]{2,3}$")
	return coderegex.MatchString(airlineCode)
}
