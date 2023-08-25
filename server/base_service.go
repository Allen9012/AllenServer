package server

import (
	imetrics "AllenServer/utils/metrics"
	"AllenServer/utils/protos"
	"context"
	"sync"
	"time"
)

// BaseService is the base class for all services.
type BaseService struct {
	Id             string
	Name           string
	DeploymentId   string
	submissionTime time.Time
	statsProcessor *imetrics.StatsProcessor // tracks and computes stats to be rendered on the /statusz page.
	traceSaver     func(spans *protos.Spans) error
	Ctx            context.Context
	mu             sync.Mutex
	Inherit        IService // 继承IService
}
