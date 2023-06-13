package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

/*
This function is responsible for generating new log events at regular intervals (every 2 seconds).

The ticker fires off an event every 2 seconds. Each time the ticker "ticks", a new log event with
the message "Hello World!" and the current timestamp is created and sent to the queueChan channel.

The function stops generating logs and returns if a signal is received on the quit channel.
*/
func getLogs(queueChan chan<- *cloudwatchlogs.InputLogEvent, quit <-chan struct{}) {
	ticker := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-ticker.C:
			// Modify the deploymentInfo to fit your use-case.
			deploymentInfo := DeploymentInfo{
				// Fill out the fields here...
			}

			// Convert deploymentInfo to JSON
			deploymentInfoJson, err := json.Marshal(deploymentInfo)
			if err != nil {
				log.Printf("Failed to marshal DeploymentInfo: %v", err)
				continue
			}

			queueChan <- &cloudwatchlogs.InputLogEvent{
				Message:   aws.String(string(deploymentInfoJson)),
				Timestamp: aws.Int64(time.Now().UnixNano() / int64(time.Millisecond)),
			}
		case <-quit:
			return
		}
	}
}
