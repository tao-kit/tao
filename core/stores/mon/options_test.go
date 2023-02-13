package mon

import (
	mopt "go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetSlowThreshold(t *testing.T) {
	assert.Equal(t, defaultSlowThreshold, slowThreshold.Load())
	SetSlowThreshold(time.Second)
	assert.Equal(t, time.Second, slowThreshold.Load())
}

func Test_defaultTimeoutOption(t *testing.T) {
	opts := mopt.Client()
	defaultTimeoutOption()(opts)
	assert.Equal(t, defaultTimeout, *opts.Timeout)
}

func TestWithTimeout(t *testing.T) {
	opts := mopt.Client()
	WithTimeout(time.Second)(opts)
	assert.Equal(t, time.Second, *opts.Timeout)
}
