package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAnnotation(t *testing.T) {
	testCases := []struct {
		testName    string
		fromCron    string
		toCron      string
		expectedErr bool
	}{
		{
			testName:    "from to not set",
			fromCron:    "",
			toCron:      "",
			expectedErr: false,
		},

		{
			testName:    "from/to with incorrect cron string",
			fromCron:    "foobar",
			toCron:      "foobar",
			expectedErr: true,
		},

		{
			testName:    "to with incorrect cron string",
			fromCron:    "",
			toCron:      "foobar",
			expectedErr: false,
		},

		{
			testName:    "from with incorrect cron string",
			fromCron:    "foobar",
			toCron:      "",
			expectedErr: false,
		},

		{
			testName:    "from set to not",
			fromCron:    "* * * * *",
			toCron:      "",
			expectedErr: false,
		},

		{
			testName:    "from and to set",
			fromCron:    "* 5 * * *",
			toCron:      "* 2 * * *",
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		_, err := newAnnotations(tc.fromCron, tc.toCron)
		if tc.expectedErr {
			assert.NotNil(t, err, tc.testName)
		} else {
			assert.Nil(t, err, tc.testName)
		}
	}
}
