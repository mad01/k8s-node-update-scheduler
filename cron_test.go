package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNewCron(t *testing.T) {
	testCases := []struct {
		testName    string
		node        *v1.Node
		expectedErr bool
	}{

		{
			testName:    "correct cron from/to strings",
			expectedErr: false,
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node0",
					Annotations: map[string]string{
						nodeAnnotationFromWindow: "* 2 * * *",
						nodeAnnotationToWindow:   "* 5 * * *",
						nodeAnnotationReboot:     "true",
					},
				},
				Spec: v1.NodeSpec{
					ProviderID: "node0",
				},
			},
		},

		{
			testName:    "incorrect from string correct to string",
			expectedErr: true,
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node0",
					Annotations: map[string]string{
						nodeAnnotationFromWindow: "notokcronstring",
						nodeAnnotationToWindow:   "* 5 * * *",
						nodeAnnotationReboot:     "true",
					},
				},
				Spec: v1.NodeSpec{
					ProviderID: "node0",
				},
			},
		},

		{
			testName:    "correct from string incorrect to string",
			expectedErr: true,
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node0",
					Annotations: map[string]string{
						nodeAnnotationFromWindow: "* 5 * * *",
						nodeAnnotationToWindow:   "foobar123",
						nodeAnnotationReboot:     "true",
					},
				},
				Spec: v1.NodeSpec{
					ProviderID: "node0",
				},
			},
		},

		{
			testName:    "correct from string incorrect to string missing one *",
			expectedErr: true,
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node0",
					Annotations: map[string]string{
						nodeAnnotationFromWindow: "* 5 * * *",
						nodeAnnotationToWindow:   "* 8 * *",
						nodeAnnotationReboot:     "true",
					},
				},
				Spec: v1.NodeSpec{
					ProviderID: "node0",
				},
			},
		},
	}

	for _, tc := range testCases {
		_, err := newMaintenanceWindow(
			tc.node.Annotations[nodeAnnotationFromWindow],
			tc.node.Annotations[nodeAnnotationToWindow],
		)
		if tc.expectedErr {
			assert.NotNil(t, err, tc.testName)
		} else {
			assert.Nil(t, err, tc.testName)
		}
	}
}
