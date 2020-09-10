// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Information about the time range of the latest available aggregated profile.
type AggregatedProfileTime struct {
	// The time period.
	Period AggregationPeriod
	// The start time.
	Start *time.Time
}

//
type AgentConfiguration struct {
	AgentParameters map[string]*string
	//
	PeriodInSeconds *int32
	//
	ShouldProfile *bool
}

//
type AgentOrchestrationConfig struct {
	//
	ProfilingEnabled *bool
}

// The description of a profiling group.
type ProfilingGroupDescription struct {
	//
	AgentOrchestrationConfig *AgentOrchestrationConfig
	// The Amazon Resource Name (ARN) identifying the profiling group.
	Arn             *string
	ComputePlatform ComputePlatform
	// The time, in milliseconds since the epoch, when the profiling group was created.
	CreatedAt *time.Time
	// The name of the profiling group.
	Name *string
	// The status of the profiling group.
	ProfilingStatus *ProfilingStatus
	// The time, in milliseconds since the epoch, when the profiling group was last
	// updated.
	UpdatedAt *time.Time
}

// Information about the profiling status.
type ProfilingStatus struct {
	// The time, in milliseconds since the epoch, when the latest agent was
	// orchestrated.
	LatestAgentOrchestratedAt *time.Time
	// The time, in milliseconds since the epoch, when the latest agent was reported..
	LatestAgentProfileReportedAt *time.Time
	// The latest aggregated profile
	LatestAggregatedProfile *AggregatedProfileTime
}

// Information about the profile time.
type ProfileTime struct {
	// The start time of the profile.
	Start *time.Time
}

type FrameMetric struct {
	FrameName    *string
	ThreadStates []*string
	Type         MetricType
}