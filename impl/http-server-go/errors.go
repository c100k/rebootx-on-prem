package main

type ServiceError struct {
	HttpStatus int
	Message    string
}

func (err *ServiceError) Error() string {
	return err.Message
}
