package input

import "github.com/esmaeel67/golang-tdd.git/calculator"

const expressionLength = 3


type Parse struct{
	engine *calculator.Engine
	validator *Validator
}

