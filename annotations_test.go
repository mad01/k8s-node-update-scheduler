package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestSetAnnotation(t *testing.T) {
	node := &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "node0",
			Annotations: map[string]string{
				"dummy": "",
			},
		},
		Spec: v1.NodeSpec{
			ProviderID: "node0",
		},
	}
	a, err := newAnnotations("* 2 * * *", "* 5 * * *")
	assert.Nil(t, err)
	node.Annotations[nodeAnnotationFromWindow] = fmt.Sprintf("%v", a.timeWindow.from)
	node.Annotations[nodeAnnotationToWindow] = fmt.Sprintf("%v", a.timeWindow.to)
	node.Annotations[nodeAnnotationReboot] = fmt.Sprintf("%v", a.reboot)
	if _, ok := node.Annotations["dummy"]; ok {
		assert.True(t, ok)
	} else {
		assert.True(t, false)
	}

}
