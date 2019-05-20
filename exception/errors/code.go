package errors

type ErrorCode uint

const (
	ErrorCodeUnknown ErrorCode = iota + 1
	ErrorCodeField
)
