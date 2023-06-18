package gorf

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type Err struct {
	Msg    string
	Er     error
	status int
}

func NewErr(msg string, status int, err error) *Err {
	var e error
	if err != nil {
		e = err
	} else {
		e = errors.New(msg)
	}
	return &Err{
		Msg:    msg,
		Er:     e,
		status: status,
	}
}

func (e *Err) Response() gin.H {
	return gin.H{
		"Message":    e.Msg,
		"Error":      e.Er.Error(),
		"StatusCode": e.status,
	}
}

func (e *Err) Error() string {
	return e.Er.Error()
}
