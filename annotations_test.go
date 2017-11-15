package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAnnotation(t *testing.T) {
	testCases := []struct {
		testName      string
		input         string
		expectedError bool
	}{

		{
			testName:      "time :01",
			input:         "01:01",
			expectedError: true,
		},

		{
			testName:      "time 01:",
			input:         "01:",
			expectedError: true,
		},

		{
			testName:      "time 01:01 AM",
			input:         "01:01 AM",
			expectedError: false,
		},

		{
			testName:      "time 01:01 am",
			input:         "01:01 am",
			expectedError: true,
		},

		{
			testName:      "time 01:01 PM",
			input:         "01:01 PM",
			expectedError: false,
		},

		{
			testName:      "time 1:01",
			input:         "1:01",
			expectedError: true,
		},

		{
			testName:      "time 1:01 AM",
			input:         "1:01 AM",
			expectedError: false,
		},

		{
			testName:      "time 1:01 PM",
			input:         "1:01 PM",
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		_, err := newAnnotations(tc.input, tc.input)
		if tc.expectedError == true {
			assert.NotNil(t, err, tc.testName)
		} else {
			assert.Nil(t, err, tc.testName)
		}
	}
}
