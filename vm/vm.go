package vm

type OpType int

const (
	OperationACC OpType = iota
	OperationJMP
	OperationNOP
)

type Instruction struct {
	OP  string
	ARG int
}
