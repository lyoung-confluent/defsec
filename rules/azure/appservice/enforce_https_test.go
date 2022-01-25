package appservice

import (
	"testing"

	"github.com/aquasecurity/defsec/provider/azure/appservice"
	"github.com/aquasecurity/defsec/state"
	"github.com/stretchr/testify/assert"
)

func TestCheckEnforceHttps(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name     string
		input    appservice.AppService
		expected bool
	}{
		{
			name:     "positive result",
			input:    appservice.AppService{},
			expected: true,
		},
		{
			name:     "negative result",
			input:    appservice.AppService{},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.Azure.AppService = test.input
			results := CheckEnforceHttps.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Rule().LongID() == CheckEnforceHttps.Rule().LongID() {
					found = true
				}
			}
			if test.expected {
				assert.True(t, found, "Rule should have been found")
			} else {
				assert.False(t, found, "Rule should not have been found")
			}
		})
	}
}