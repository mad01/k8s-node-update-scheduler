package main

import "fmt"

const (
	nodeAnnotationReboot     = "k8s.node.terminator.reboot" // true as string
	nodeAnnotationFromWindow = "k8s.node.terminator.fromTimeWindow"
	nodeAnnotationToWindow   = "k8s.node.terminator.toTimeWindow"
)

func newAnnotations(fromCronTime, toCronTime string) (*Annotations, error) {
	a := Annotations{
		reboot:     "false",
		timeWindow: nil,
	}
	if fromCronTime != "" && toCronTime != "" {
		m, err := newMaintenanceWindow(fromCronTime, toCronTime)
		if err != nil {
			return nil, fmt.Errorf("failed to create new annotation: %v", err.Error())
		}
		a.timeWindow = m
	}
	return &a, nil
}

// Annotations all annotaitons to add to node
type Annotations struct {
	reboot     string // true or false as string
	timeWindow *MaintenanceWindow
}

// Get annotation map
func (a *Annotations) Get() map[string]string {
	m := map[string]string{
		nodeAnnotationReboot:     a.reboot,
		nodeAnnotationFromWindow: fmt.Sprintf("%v", a.timeWindow.from),
		nodeAnnotationToWindow:   fmt.Sprintf("%v", a.timeWindow.to),
	}
	return m
}
