package utils

type ServiceError struct {
	HTTPStatus int
	Message    string
}

func (err *ServiceError) Error() string {
	return err.Message
}
