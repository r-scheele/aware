package main

import (
	"os"
	"os/signal"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func init() {
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(awsRegion),
		},
	}))

	cwl = cloudwatchlogs.New(sess)
	err := ensureLogGroupExists(logGroupName)
	if err != nil {
		panic(err)
	}
}

func main() {
	queueChan := make(chan *cloudwatchlogs.InputLogEvent, 1000)
	quit := make(chan struct{})

	go getLogs(queueChan, quit)

	go processQueue(queueChan, quit)

	// Wait for a user interrupt, then quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	// Block until we receive an interrupt signal, then close the quit channel
	<-sig
	close(quit)
}
