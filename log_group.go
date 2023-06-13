package main

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

/*
ensureLogGroupExists first checks if the log group exists, if it doesn't it will create one.
*/
func ensureLogGroupExists(name string) error {
	resp, err := cwl.DescribeLogGroups(&cloudwatchlogs.DescribeLogGroupsInput{})
	if err != nil {
		return err
	}

	for _, logGroup := range resp.LogGroups {
		if *logGroup.LogGroupName == name {
			return nil
		}
	}

	_, err = cwl.CreateLogGroup(&cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: &name,
	})
	if err != nil {
		return err
	}

	_, err = cwl.PutRetentionPolicy(&cloudwatchlogs.PutRetentionPolicyInput{
		RetentionInDays: &retentionDays,
		LogGroupName:    &name,
	})

	return err
}
