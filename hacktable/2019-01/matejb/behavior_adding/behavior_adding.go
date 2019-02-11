package main

import (
	"fmt"
	"strconv"
	"time"
)

type TemporaryError interface {
	Retry() time.Time
}

func main() {
	handleOperationErr := func(opName string, err error) {
		fmt.Printf("operation %s failed with error: %s", opName, err)
		if tmpErr, ok := err.(TemporaryError); ok {
			fmt.Printf("; will retry after %v", tmpErr.Retry().Sub(time.Now()))
		}
		fmt.Println("")
	}

	handleOperationErr("A", operatioA())
	handleOperationErr("B", operatioB())
}

func operatioA() error {
	return appError{code: 100}
}

func operatioB() error {
	return appError{code: 200}.asTemporary(time.Now().Add(time.Minute))
}

type appError struct {
	code int
}

func (ae appError) Error() string {
	return strconv.Itoa(ae.code)
}

func (ae appError) asTemporary(retry time.Time) error {
	return temporaryAppError{appError: ae, retry: retry}
}

type temporaryAppError struct {
	appError
	retry time.Time
}

func (err temporaryAppError) Retry() time.Time {
	return err.retry
}
