package error_test

import (
	"main/utils/error"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

// ErrorSuite is a test suite for the error package
type ErrorSuite struct {
	suite.Suite
}

// TestErrorToString tests the ToString method of the Error type
func (suite *ErrorSuite) TestErrorToString() {
	e := error.NewError(http.StatusNotFound, "Resource not found")
	suite.Equal("Resource not found", e.ToString(), "Error message should match")
}

// TestToError tests the ToError method
func (suite *ErrorSuite) TestToError() {
    customErr := error.NewError(http.StatusBadRequest, "custom error")
    suite.NotNil(customErr)

    err := customErr.ToError()
    suite.NotNil(suite.T(), err)
    suite.Equal("custom error", err.Error(), "Error message should match")
}

// TestNewError tests the creation of a new generic error
func (suite *ErrorSuite) TestNewError() {
	code := http.StatusBadRequest
	msg := "Invalid request"

	e := error.NewError(code, msg)

	suite.NotNil(e, "Error should not be nil")
	suite.Equal(code, e.Code, "Error code should match")
	suite.Equal(msg, e.Message, "Error message should match")
}

// TestNewNotFoundError tests the creation of a not found error
func (suite *ErrorSuite) TestNewNotFoundError() {
	msg := "Resource not found"
	e := error.NewNotFoundError(msg)

	suite.NotNil(e, "Error should not be nil")
	suite.Equal(http.StatusNotFound, e.Code, "Error code should be 404")
	suite.Equal(msg, e.Message, "Error message should match")
}

// TestNewUnexpectedError tests the creation of an internal server error
func (suite *ErrorSuite) TestNewUnexpectedError() {
	msg := "An unexpected error occurred"
	e := error.NewUnexpectedError(msg)

	suite.NotNil(e, "Error should not be nil")
	suite.Equal(http.StatusInternalServerError, e.Code, "Error code should be 500")
	suite.Equal(msg, e.Message, "Error message should match")
}

// TestNewValidationError tests the creation of a validation error
func (suite *ErrorSuite) TestNewValidationError() {
	msg := "Validation failed"
	e := error.NewValidationError(msg)

	suite.NotNil(e, "Error should not be nil")
	suite.Equal(http.StatusUnprocessableEntity, e.Code, "Error code should be 422")
	suite.Equal(msg, e.Message, "Error message should match")
}

// TestErrorSuite initializes the test suite
func TestErrorSuite(t *testing.T) {
	suite.Run(t, new(ErrorSuite))
}