# structerror

A error with code and message.

- easy to compare with the code.
- output as struct string like json.

```
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
```

## CodeError

An impletement of struct error.

```
type CodeError struct {
	Code   int
	Status string
	Msg    string
	Causes []error
}
```


