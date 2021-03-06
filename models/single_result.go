package models

// StatusError used to represent errored query or operation
const StatusError = "error"

// StatusSuccess used to represent success query or operation
const StatusSuccess = "success"

// SingleResult used to wrap single record result from query or create operation
type SingleResult struct {
	item     Modellable
	err      error
	status   string
	notFound bool
}

// MakeSingleResult used to instantiate single result
func MakeSingleResult() *SingleResult {
	result := &SingleResult{}

	return result
}

// StatusError used to check if operation is errored
func (result *SingleResult) StatusError() bool {
	return result.status == StatusError
}

// StatusSuccess used to check if operation is success
func (result *SingleResult) StatusSuccess() bool {
	return result.status == StatusSuccess
}

// Status used to retrieve result's status
func (result *SingleResult) Status() string {
	return result.status
}

func (result *SingleResult) Error() error {
	return result.err
}

// Set used to set result item, error and status
func (result *SingleResult) Set(item Modellable, err error, notFound bool) {
	result.item = item
	result.err = err
	result.notFound = notFound

	if err != nil {
		result.status = StatusError
	} else {
		result.status = StatusSuccess
	}
}

// ErrorMessage used to retrieve result's error message
func (result *SingleResult) ErrorMessage() string {
	if result.err == nil {
		return ""
	}

	return result.err.Error()
}

// Item used to retrieve result's item
func (result *SingleResult) Item() Modellable {
	return result.item
}

func (result *SingleResult) NotFound() bool {
	return result.notFound
}
