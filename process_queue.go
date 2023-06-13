package main

import (
	"time"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

/*
This function is responsible for processing the queue of log events at regular intervals (every 5 seconds).

The ticker fires off an event every 5 seconds. Each time the ticker "ticks", if there are log events in the
logQueue, they are sent to CloudWatch Logs using the sendLogEvents function. After the log events are sent,
the logQueue is cleared.

If a log event is received on the queueChan channel, it is added to the logQueue.

The function stops processing the queue and returns if a signal is received on the quit channel.
*/
func processQueue(queueChan chan *cloudwatchlogs.InputLogEvent, quit chan struct{}) {
	var logQueue []*cloudwatchlogs.InputLogEvent
	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case logEvent := <-queueChan:
			logQueue = append(logQueue, logEvent)
		case <-ticker.C:
			if len(logQueue) > 0 {
				sendLogEvents(logQueue)
				logQueue = []*cloudwatchlogs.InputLogEvent{}
			}
		case <-quit:
			// Process remaining log events before exiting
			if len(logQueue) > 0 {
				sendLogEvents(logQueue)
			}
			return
		}
	}
}
