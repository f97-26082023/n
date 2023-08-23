package api

import (
	"github.com/f97/n/pkg/core"
	"github.com/f97/n/pkg/errs"
	"github.com/f97/n/pkg/settings"
)

// HealthsApi represents health api
type HealthsApi struct{}

// Initialize a healths api singleton instance
var (
	Healths = &HealthsApi{}
)

// HealthStatusHandler returns the health status of current service
func (a *HealthsApi) HealthStatusHandler(c *core.Context) (interface{}, *errs.Error) {
	result := make(map[string]string)

	result["version"] = settings.Version
	result["commit"] = settings.CommitHash
	result["status"] = "ok"

	return result, nil
}
