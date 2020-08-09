package envvarstore

import (
	"testing"

	"github.com/rbonnat/codecademy/configuration"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {

	tests := []struct {
		name             string
		mockVarName      string
		mockVarValue     string
		varName          string
		expectedVarValue string
	}{
		{
			"Successful",
			"testName",
			"testValue",
			"testName",
			"testValue",
		},
	}

	for _, test := range tests {
		mockVarStore := configuration.MockVarStore{}
		mockVarStore.On("Get", test.mockVarName).Return(test.mockVarValue)
		v := mockVarStore.Get(test.varName)
		assert.Equal(t, test.expectedVarValue, v)
	}

}
