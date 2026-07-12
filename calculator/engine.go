package calculator

import (
	"fmt"

	"github.com/esmaeel67/golang-tdd.git/format"
)

type Operation struct {
	Expression string
	Operator   string
	Operands   []float64
}

type Engine struct {
	expectedLength  int
	validOperations map[string]func(x, y float64) float64
}

func NewEngine() *Engine {
	e := Engine{
		expectedLength:  2,
		validOperations: make(map[string]func(x float64, y float64) float64),
	}
	e.validOperations["+"] = e.Add

	return &e
}

func (e *Engine) GetNumOperands() int {
	return e.expectedLength
}

func (e *Engine) GetValidOperators() []string {
	var ops []string
	for o := range e.validOperations {
		ops = append(ops, o)
	}
	return ops
}

func (e *Engine) ProcessOperation(operation Operation) (*string, error) {
	f, ok := e.validOperations[operation.Operator]
	if !ok {
		err := fmt.Errorf("no operation for operator %s found", operation.Operator)
		return nil, format.Error(operation.Expression, err)
	}
	res := f(operation.Operands[0], operation.Operands[1])
	fres := format.Result(operation.Expression, res)
	return &fres, nil
}

func (e *Engine) Add(x, y float64) float64 {
	return x + y
}
