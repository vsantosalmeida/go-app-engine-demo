package person

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePersonBirthDate(t *testing.T) {
	t.Run("When a birth date is in correct layout must return a valid time and no err", func(t *testing.T) {
		d, err := parsePersonBirthDate("1998-06-12")
		assert.NoError(t, err)
		assert.NotZero(t, d, "without error d must has a valid value")
	})

	t.Run("When a birth date isn't in correct layout must return a err", func(t *testing.T) {
		d, err := parsePersonBirthDate("2000-31-05")
		assert.Error(t, err)
		assert.Zerof(t, d, "after error d must be zero value")
	})
}
