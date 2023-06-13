# Aware

## Overview

Aware generates log events about a deployment send them to AWS CloudWatch Logs. 

The application has a primary log group under which log streams are created. Each log stream contains log events. Log events are generated at regular intervals (every 2 seconds) and are processed in batches (every 5 seconds). The application keeps running and processing logs until a user interrupt signal (CTRL+C) is received.

## Main Packages

The following packages are imported for the operation of this program:

`os and os/signal`: To catch the user interrupt signal. \
`time`: To schedule the generation and processing of log events. \
`log`: To log errors that might occur during the execution of the program. \
`encoding/json`: To encode log data as JSON. \
`github.com/aws/aws-sdk-go/aws`: To interact with the AWS SDK. \
`github.com/aws/aws-sdk-go/aws/session`: To create a new AWS SDK session. \
`github.com/aws/aws-sdk-go/service/cloudwatchlogs`: To interact with AWS CloudWatch Logs. \
`github.com/google/uuid`: To generate unique identifiers for log streams. \

## Global Variables

The following global variables are defined in the application:

`sess`: An AWS SDK session object. \
`cwl`: A CloudWatch Logs service client. \
`logGroupName`: The name of the CloudWatch Logs group. \
`logStreamName`: The name of the CloudWatch Logs stream. \
`sequenceToken`: A token used for ordering log events.  \
`retentionDays`: The number of days to retain the logs. \
`awsRegion`: The AWS region where the logs should be stored. \

## Structs

Two structs are defined in the application:

`DeploymentInfo`: Encapsulates information about a deployment process. Each field represents a different aspect of the deployment, such as the time the deployment started, the version of the application that was deployed, and the status of the deployment.
`TestResultInfo`: Encapsulates information about the test results in a deployment. Each field represents a different aspect of the test results, such as the number of tests passed and failed, the code coverage, and other metrics related to the tests.

## Main Functions

The following main functions are defined in the application:

`init()`: This function is automatically called before main(). It initializes an AWS SDK session and a CloudWatch Logs service client. It also ensures that the specified log group exists. \
`main()`: This function starts two goroutines, getLogs() and processQueue(), and waits for a user interrupt signal. When a user interrupt signal is received, it sends a quit signal to the goroutines, which then stop their operations. \
`processQueue(queueChan chan *cloudwatchlogs.InputLogEvent, quit chan struct{})`: This function processes the queue of log events at regular intervals. If there are log events in the queue, they are sent to CloudWatch Logs. The function stops processing the queue if a quit signal is received. \
`getLogs(queueChan chan<- *cloudwatchlogs.InputLogEvent, quit <-chan struct{})`: This function generates new log events at regular intervals. Each log event is sent to the queueChan channel. The function stops generating logs if a quit signal is received. \
`ensureLogGroupExists(name string) error`: This function checks if a log group exists. If it doesn't, the function creates the log group. \
`createLogStream() error`: This function creates a new log stream with a random UUID as its