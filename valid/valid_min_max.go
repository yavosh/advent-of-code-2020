package valid

import "strconv"

type MinMaxChecker struct {
	Min int
	Max int
}

func (c *MinMaxChecker) Check(value string) bool {
	if value == "" {
		return false
	}

	intVal, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	if intVal > c.Max || intVal < c.Min {
		return false
	}

	return true
}
