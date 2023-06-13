package main

import (
	"log"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

/*
The sendLogEvents function is responsible for sending a batch of log events to AWS CloudWatch Logs.

It takes a slice of log events as input and creates a PutLogEventsInput object with the log events and the
name of the log group.

If sequenceToken is not set (i.e., it's an empty string), it means that a new log stream needs to be created
in the log group. The createLogStream function is called to create the log stream, and any errors are handled
by panicking.

If sequenceToken is set, it is added to the PutLogEventsInput object.

The name of the log stream is then added to the PutLogEventsInput object.

The PutLogEvents method of the CloudWatch Logs client is then called with the PutLogEventsInput object. Any
errors that occur during the call are logged.

If the call to PutLogEvents is successful, the sequenceToken is updated with the NextSequenceToken from the
response. This sequenceToken will be used for the next batch of log events that is sent to CloudWatch Logs.
*/
func sendLogEvents(logQueue []*cloudwatchlogs.InputLogEvent) {
	input := cloudwatchlogs.PutLogEventsInput{
		LogEvents:    logQueue,
		LogGroupName: &logGroupName,
	}

	if sequenceToken == "" {
		err := createLogStream()
		if err != nil {
			panic(err)
		}
	} else {
		input = *input.SetSequenceToken(sequenceToken)
	}

	input = *input.SetLogStreamName(logStreamName)

	resp, err := cwl.PutLogEvents(&input)
	if err != nil {
		log.Println(err)
	}

	if resp != nil {
		sequenceToken = *resp.NextSequenceToken
	}
}
