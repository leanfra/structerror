package structerror

import (
	"encoding/json"
	"fmt"
)

const (
	// CodeUnknown is unknown error code
	CodeUnknown = 1099
	// StatusUnknown is unknown error code status
	StatusUnknown = "Unknown"
)

// CodeErrorFactory is to generate new CodeError
type CodeErrorFactory struct {
	codeMap map[int]string
}

// CodeError is error with code
type CodeError struct {
	Code   int
	Status string
	Msg    string
	Causes []error
}

// NewCodeErrorFactory impletment a CodeErrorFactory
func NewCodeErrorFactory(codeMap map[int]string) CodeErrorFactory {
	return CodeErrorFactory{
		codeMap: codeMap,
	}
}

// NewError generate a new CodeError
func (f CodeErrorFactory) Error(code int, msg string, causes ...error) CodeError {
	st, ok := f.codeMap[code]
	if !ok {
		st = StatusUnknown
	}

	return CodeError{
		Code:   code,
		Status: st,
		Msg:    msg,
		Causes: causes,
	}
}

// Error impletment Error interface
func (ce CodeError) Error() string {
	return fmt.Sprintf("%s:%s", ce.Status, ce.Msg)
}

// Code impletment Code interface
func (ce CodeError) ErrorCode() int {
	return ce.Code
}

// JSON impletment StructError interface
func (ce CodeError) JSON() (string, error) {
	d, e := json.Marshal(ce)
	if e != nil {
		return "", e
	}
	return string(d), nil
}

// MarshalJSON impletement json.Marshal interface
func (ce CodeError) MarshalJSON() ([]byte, error) {
	type strCodeError struct {
		Code   int      `json:"code"`
		Status string   `json:"status"`
		Msg    string   `json:"msg"`
		Causes []string `json:"causes"`
	}

	strCauses := []string{}
	for _, e := range ce.Causes {
		strCauses = append(strCauses, e.Error())
	}

	sj := strCodeError{
		Code:   ce.Code,
		Status: ce.Status,
		Msg:    ce.Msg,
		Causes: strCauses,
	}
	return json.Marshal(sj)
}
