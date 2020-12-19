package valid

type OneOfChecker struct {
	Expected map[string]bool
}

func NewOneOfChecker(values ...string) *OneOfChecker {
	expected := make(map[string]bool)
	for _, val := range values {
		expected[val] = true
	}

	return &OneOfChecker{Expected: expected}
}

func (c *OneOfChecker) Check(value string) bool {
	if value == "" {
		return false
	}

	return c.Expected[value]
}
