package error

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

type StatusError struct {
	Err     error
	Status  int
	Title   string
	Message string
	File    string
	Line    int
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("Status %d at %s : %s", e.Status, e.FileLine(), e.Err)
}

func (e *StatusError) String() string {
	return fmt.Sprintf("Status %d at %s : %s %s %s", e.Status, e.FileLine(), e.Title, e.Message, e.Err)
}

func (e *StatusError) FileLine() string {
	parts := strings.Split(e.File, "/")
	f := strings.Join(parts[len(parts)-4:len(parts)], "/")
	return fmt.Sprintf("%s:%d", f, e.Line)
}

func (e *StatusError) setupFromArgs(args ...string) *StatusError {
	if e.Err == nil {
		e.Err = fmt.Errorf("Error:%d", e.Status)
	}
	if len(args) > 0 {
		e.Title = args[0]
	}
	if len(args) > 1 {
		e.Message = args[1]
	}
	return e
}

func NotFoundError(e error, args ...string) *StatusError {
	err := Error(e, http.StatusNotFound, "Not Found", "Sorry, the page you're looking for couldn't be found.")
	return err.setupFromArgs(args...)
}

func InternalError(e error, args ...string) *StatusError {
	err := Error(e, http.StatusInternalServerError, "Server Error", "Sorry, something went wrong, please let us know.")
	return err.setupFromArgs(args...)
}

func BadRequestErr(e error, args ...string) *StatusError {
	err := Error(e, http.StatusBadRequest, "Bad request", "bad request")
	return err.setupFromArgs(args...)
}

func Error(e error, s int, t string, m string) *StatusError {
	_, f, l, _ := runtime.Caller(2)
	err := &StatusError{
		Status:  s,
		Err:     e,
		Title:   t,
		Message: m,
		File:    f,
		Line:    l,
	}
	return err
}

func HandlerErrorGin(c *gin.Context, err error) {
	c.Next()
	c.Error(err)
}
