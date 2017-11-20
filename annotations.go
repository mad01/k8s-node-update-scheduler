package main

import (
	"fmt"

	"github.com/mad01/node-terminator/pkg/annotations"
	"github.com/mad01/node-terminator/pkg/window"
)

func newAnnotations(fromTime, toTime string) (*Annotations, error) {
	a := Annotations{
		reboot:     "false",
		timeWindow: nil,
	}
	if fromTime != "" && toTime != "" {
		m, err := window.NewMaintenanceWindow(fromTime, toTime)
		if err != nil {
			return nil, fmt.Errorf("failed to create new annotation: %v", err.Error())
		}
		a.timeWindow = m
	} else {
		a.timeWindow = &window.MaintenanceWindow{}
	}

	return &a, nil
}

// Annotations all annotaitons to add to node
type Annotations struct {
	reboot     string // true or false as string
	timeWindow *window.MaintenanceWindow
}

// Get annotation map
func (a *Annotations) Get() map[string]string {
	m := map[string]string{
		annotations.NodeAnnotationReboot:     a.reboot,
		annotations.NodeAnnotationFromWindow: fmt.Sprintf("%v", a.timeWindow.From()),
		annotations.NodeAnnotationToWindow:   fmt.Sprintf("%v", a.timeWindow.To()),
	}
	return m
}
