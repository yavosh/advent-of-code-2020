package valid

import (
	"strconv"
)

type HeightChecker struct {
	MinCm   int
	MaxCm   int
	MinInch int
	MaxInch int
}

func (c *HeightChecker) Check(value string) bool {
	if value == "" {
		return false
	}

	if len(value) < 3 {
		return false
	}

	unit := value[len(value)-2:]
	unitValue := value[:len(value)-2]

	intVal, err := strconv.Atoi(unitValue)
	if err != nil {
		return false
	}

	if unit == "cm" {
		if intVal > c.MaxCm || intVal < c.MinCm {
			return false
		}

		return true
	}

	if unit == "in" {
		if intVal > c.MaxInch || intVal < c.MinInch {
			return false
		}

		return true
	}

	return false
}
