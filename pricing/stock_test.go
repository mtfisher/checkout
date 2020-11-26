package pricing_test

import (
	"testing"

	"github.com/github.com/mtfisher/checkout/pricing"
	"github.com/stretchr/testify/assert"
)

func TestStock_CalculatePrice(t *testing.T) {
	s := pricing.Stock{UnitPrice: 20}
	assert.Exactly(t, 40, s.CalculatePrice(2))
}
