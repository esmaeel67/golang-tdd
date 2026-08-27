package input_test

import (
	"fmt"
	"testing"

	"github.com/esmaeel67/golang-tdd.git/calculator"
	"github.com/esmaeel67/golang-tdd.git/input"
	"github.com/esmaeel67/golang-tdd.git/mocks"
	"github.com/stretchr/testify/assert"
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
		assert.Contains(t, *result, expectedResult)
		assert.Contains(t, *result, expr)
		validator.AssertExpectations(t)
		engine.AssertExpectations(t)
	})

	t.Run("invalid operator", func(t *testing.T) {
		// Arrange
		expr := "2 % 3"
		operator := "%"
		operands := []float64{2.0, 3.0}
		expectedErrMsg := "bad operator"
		engine := mocks.NewOperationProcessor(t)
		validator := mocks.NewValidationHelper(t)
		parser := input.NewParser(engine, validator)
		validator.On("CheckInput", operator, operands).Return(fmt.Errorf("%s", expectedErrMsg)).Once()

		// Act
		result, err := parser.ProcessExpression(expr)

		// Assert
		require.NotNil(t, err)
		require.Nil(t, result)
		assert.Contains(t, err.Error(), expectedErrMsg)
		validator.AssertExpectations(t)
	})

	t.Run("invalid right operand", func(t *testing.T) {
		expr := "2 + abc"
		expectedErrMsg := "unable to process"
		engine := mocks.NewOperationProcessor(t)
		validator := mocks.NewValidationHelper(t)
		parser := input.NewParser(engine, validator)

		// Act
		result, err := parser.ProcessExpression(expr)

		// Assert
		require.NotNil(t, err)
		require.Nil(t, result)
		assert.Contains(t, err.Error(), expr)
		assert.Contains(t, err.Error(), expectedErrMsg)
		validator.AssertExpectations(t)
	})

	t.Run("invalid left operand", func(t *testing.T) {
		// Arrange
		expr := "abc + 2"
		expectedErrMsg := "unable to process"
		engine := mocks.NewOperationProcessor(t)
		validator := mocks.NewValidationHelper(t)
		parser := input.NewParser(engine, validator)

		// Act
		result, err := parser.ProcessExpression(expr)

		// Asset
		require.Error(t, err)
		require.Nil(t, result)
		assert.Contains(t, err.Error(), expr)
		assert.Contains(t, err.Error(), expectedErrMsg)
		validator.AssertExpectations(t)
	})

	t.Run("invalid expression length", func(t *testing.T) {
		// Arrange
		expr := "2 + 3 + 5"
		expectedErrMsg := "incorrect expression length"
		engine := mocks.NewOperationProcessor(t)
		validator := mocks.NewValidationHelper(t)
		parser := input.NewParser(engine, validator)

		// Act
		result, err := parser.ProcessExpression(expr)

		// Assert
		require.Error(t, err)
		require.Nil(t, result)
		assert.Contains(t, err.Error(), expr)
		assert.Contains(t, err.Error(), expectedErrMsg)
		validator.AssertExpectations(t)
	})
	// t.Skip("not implemented yet, will implement in chapter03")

}
