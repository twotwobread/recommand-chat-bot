package assertion

import (
	"reflect"
	"testing"

	"recommand-chat-bot/domain"

	"github.com/stretchr/testify/assert"
)

func AssertMovieInputFields(t *testing.T, expected domain.Movie, actual domain.Movie) {
	vExpected := reflect.ValueOf(expected)
	vActual := reflect.ValueOf(actual)

	// Check if the expected or actual values are pointers, and dereference if needed
	if vExpected.Kind() == reflect.Ptr {
		vExpected = vExpected.Elem()
	}
	if vActual.Kind() == reflect.Ptr {
		vActual = vActual.Elem()
	}

	// Iterate over the struct fields and assert equality
	for i := 0; i < vExpected.NumField(); i++ {
		fieldName := vExpected.Type().Field(i).Name

		// Skip auto-generated fields like ID, CreatedAt, UpdatedAt
		if fieldName == "ID" || fieldName == "CreatedAt" || fieldName == "UpdatedAt" {
			continue
		}

		expectedValue := vExpected.Field(i).Interface()
		actualValue := vActual.Field(i).Interface()

		// Assert that the values are equal
		assert.Equal(t, expectedValue, actualValue)
	}
}
