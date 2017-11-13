package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

func newMaintenanceWindow(from, to string) (*MaintenanceWindow, error) {
	exprStart, err := cronexpr.Parse(from)
	if err != nil {
		return nil, fmt.Errorf("failed to Parse cronexpr %v", err.Error())
	}
	exprStop, err := cronexpr.Parse(to)
	if err != nil {
		return nil, fmt.Errorf("failed to Parse cronexpr %v", err.Error())
	}

	windowStartTime := exprStart.Next(time.Now())
	windowStopTime := exprStop.Next(time.Now())

	m := MaintenanceWindow{
		fromCron: from,
		toCron:   to,
		from:     &windowStartTime,
		to:       &windowStopTime,
	}
	return &m, nil

}

// MaintenanceWindow info
type MaintenanceWindow struct {
	fromCron string
	toCron   string
	from     *time.Time
	to       *time.Time
}
