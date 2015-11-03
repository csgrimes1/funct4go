package lists_test
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/csgrimes1/funct4go/lists"
)

func TestTrueMaybe(t *testing.T) {
	result := lists.CreateBooleanResult(true)
	count := result.Count()
	assert.Equal(t, 1, count)
}

func TestFalseMaybe(t *testing.T) {
	result := lists.CreateBooleanResult(false)
	count := result.Count()
	assert.Equal(t, 0, count)
}

func TestOptionalResult(t *testing.T) {
	result := lists.CreateOptionalResult("hello whirreled")
	count := result.Count()
	assert.Equal(t, 1, count)
}

func TestOptionalEmpty(t *testing.T) {
	result := lists.CreateOptionalEmpty()
	count := result.Count()
	assert.Equal(t, 0, count)
}

