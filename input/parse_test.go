package input_test

import (
	"testing"

	"github.com/esmaeel67/golang-tdd.git/calculator"
	"github.com/esmaeel67/golang-tdd.git/input"
	"github.com/esmaeel67/golang-tdd.git/mocks"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {

	t.Run("valid input", func(t *testing.T) {
		// Arrange
		expr := "2 + 3"
		operator := "+"
		operands := []float64{2.0, 3.0}
		expectedResult := "2 + 3 = 5.5"
		engine := mocks.NewOperationProcessor(t)
		validator := mocks.NewValidationHelper(t)
		parser := input.NewParser(engine, validator)

		validator.On("CheckInput", operator, operands).Return(nil).Once()
		engine.On("ProcessOperation", calculator.Operation{
			Expression: expr,
			Operator:   operator,
			Operands:   operands,
		}).Return(&expectedResult, nil).Once()

		// Act
		result, err := parser.ProcessExpression(expr)

		// Assert
		require.Nil(t, err)
		require.NotNil(t, result)
	})

	t.Skip("not implemented yet, will implement in chapter03")

}
