package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	// assert.EqualValues(t, 1, len(err.causes))
	// assert.EqualValues(t, "database error", err.causes[0])

}
func TestNewBadRequestError(t *testing.T) {

}
func TestNewNotFoundError(t *testing.T) {

}
func TestNewError(t *testing.T) {

}
