package main

import (
	"time"

	"github.com/gorhill/cronexpr"
)

func newMaintenanceWindow(from, to string) *MaintenanceWindow {
	exprStart := cronexpr.MustParse(from)
	exprStop := cronexpr.MustParse(to)

	windowStartTime := exprStart.Next(time.Now())
	windowStopTime := exprStop.Next(time.Now())

	m := MaintenanceWindow{
		from: &windowStartTime,
		to:   &windowStopTime,
	}
	return &m

}

// MaintenanceWindow info
type MaintenanceWindow struct {
	from *time.Time
	to   *time.Time
}
