package rest

const (
	defaultErrorCode = -1
)

type VmwareFusionError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func FromRawError(err error) VmwareFusionError {
	return VmwareFusionError{
		Code:    defaultErrorCode,
		Message: err.Error(),
	}
}
