package testing

import (
	"github.com/stretchr/testify/assert"
	"payment-payments-api/pkg/optimizer"
	"testing"
)

var optimizerEngine = optimizer.NewOptimizer("/home/abraham/falabella/python/altiro-optimizer/weboptimizer.py", "python3.8")

func TestInitOptimizerTest(t *testing.T) {
	assert := assert.New(t)

	assert.NotNil(optimizerEngine)
}
