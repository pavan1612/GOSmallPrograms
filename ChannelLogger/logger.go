package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	timeOccured time.Time
	severity    string
	message     string
}

var logCh = make(chan logEntry)
var doneCh = make(chan struct{})


func main() {
	go logger()

	logCh <- logEntry{time.Now(), logInfo, "App has started"}
	logCh <- logEntry{time.Now(), logWarning, "App has a Warning"}

	doneCh <- struct{}{}


}

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v : Severity : %v - %v\n", entry.timeOccured.Format("15:04:05.000"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}

}
