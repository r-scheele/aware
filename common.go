package main

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

var (
	sess          *session.Session
	cwl           *cloudwatchlogs.CloudWatchLogs
	logGroupName  = "MyApplication" // Container for log streams
	logStreamName = "MyLogStream"   // Container for log events - Log category
	sequenceToken string            // Individually identified log events
	retentionDays = int64(7)
	awsRegion     = "us-east-1"
)

// DeploymentInfo encapsulates information about a deployment process.
type DeploymentInfo struct {
	ApplicationVersion    string         // Version of the application that was deployed.
	Initiator             string         // User or system that initiated the deployment.
	Environment           string         // Environment where the deployment was made (e.g., dev, staging, prod).
	DeploymentMethod      string         // Method of deployment (e.g., manual, automatic, rollback).
	ChangeList            []string       // List of changes included in the deployment.
	Status                string         // Status of the deployment (successful, failed, etc.) The current state of the deployment as a whole..
	Errors                []string       // List of errors occurred during the deployment, if any.
	DeploymentDuration    time.Duration  // Duration of the deployment process. This can be calculated from start_time and end_time.
	StartTime             time.Time      // Time when the deployment process started.
	EndTime               time.Time      // Time when the deployment process ended.
	InfrastructureChanges []string       // Any changes made to the underlying infrastructure.
	AffectedServices      []string       // Services affected by the deployment.
	RollbackInformation   string         // Information for rollback in case of failure.
	DeploymentToolVersion string         // Version of the deployment tool used.
	TestResults           TestResultInfo // Test results if tests were run during the deployment.
}

// TestResultInfo encapsulates information about the test results in a deployment.
type TestResultInfo struct {
	TestsPassed  int                    // Number of tests passed.
	TestsFailed  int                    // Number of tests failed.
	CodeCoverage float64                // Code coverage in the tests.
	OtherMetrics map[string]interface{} // Other metrics related to the tests.
}
