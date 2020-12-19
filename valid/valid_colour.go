package valid

import "regexp"

type ColourChecker struct {
}

func (c *ColourChecker) Check(value string) bool {

	if value == "" {
		return false
	}

	found, err := regexp.MatchString("^#[a-f0-9]{6}$", value)
	if err != nil {
		return false
	}

	return found
}
