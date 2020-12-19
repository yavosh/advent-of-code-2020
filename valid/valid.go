package valid

type Checker interface {
	Check(v string) bool
}
