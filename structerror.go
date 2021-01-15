package structerror

// StructError interface
type StructError interface {
	error
	ErrorCode() int
}

// StructJsonError interface
type StructJsonError interface {
	StructError
	JSON() (string, error)
}
