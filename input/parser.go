package input

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/esmaeel67/golang-tdd.git/calculator"
	"github.com/esmaeel67/golang-tdd.git/format"
)

const expressionLength = 3

type Parser struct {
	engine    *calculator.Engine
	validator *Validator
}

func NewParser(op *calculator.Engine, v *Validator) *Parser {
	return &Parser{
		engine:    op,
		validator: v,
	}
}

func (p *Parser) ProcessExpression(expr string) (*string, error) {
	operation, err := p.getOperation(expr)
	if err != nil {
		return nil, format.Error(expr, err)
	}
	return p.engine.ProcessOperation(*operation)
}

func (p *Parser) getOperation(expr string) (*calculator.Operation, error) {
	ops := strings.Fields(expr)
	if len(ops) != expressionLength {
		return nil, fmt.Errorf("incorrect expression length:got %d, want %d", len(ops), expressionLength)
	}
	leftOp, err := strconv.ParseFloat(ops[0], 64)
	if err != nil {
		return nil, fmt.Errorf("unable to process expression: %v", err)
	}
	rightOp, err := strconv.ParseFloat(ops[2], 64)
	if err != nil {
		return nil, fmt.Errorf("unable to process expression: %v", err)
	}
	operator := ops[1]
	operands := []float64{leftOp, rightOp}
	if err := p.validator.CheckInput(operator, operands); err != nil {
		return nil, err
	}

	return &calculator.Operation{
		Expression: expr,
		Operator:   operator,
		Operands:   operands,
	}, nil
}
