package expiredates_test

import (
	"fmt"
	"testing"

	"github.com/hsmtkk/mf-vix-strategy/expiredates"
	"github.com/stretchr/testify/assert"
)

func TestGetExpireDates(t *testing.T) {
	dates, err := expiredates.GetExpireDates()
	assert.Nil(t, err)
	assert.NotNil(t, dates)
	fmt.Printf("%v\n", dates)
}
