package main

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/google/uuid"
)

// createLogStream will make a new logStream with a random uuid as its name.
func createLogStream() error {
	name := uuid.New().String()

	_, err := cwl.CreateLogStream(&cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  &logGroupName,
		LogStreamName: &name,
	})

	logStreamName = name

	return err
}
