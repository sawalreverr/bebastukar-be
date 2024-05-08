package pkg

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/bebastukar-be/internal/helper"
)

var (
	ErrRecordNotFound      = errors.New("record not found")
	ErrDataAlreadyExist    = errors.New("data already exist")
	ErrStatusInternalError = errors.New("internal server error")

	ErrNoPrivilege = errors.New("no permission to doing this task")

	// discussion error
	ErrDiscussionNotFound   = errors.New("discussion not found")
	ErrCommentNotFound      = errors.New("comment not found")
	ErrReplyCommentNotFound = errors.New("reply comment not found")
)

func DiscussionErrorHelper(c echo.Context, err error) error {
	var errorResponse string

	if errors.Is(err, ErrDiscussionNotFound) {
		errorResponse = err.Error()
	}

	if errors.Is(err, ErrCommentNotFound) {
		errorResponse = err.Error()
	}

	if errors.Is(err, ErrReplyCommentNotFound) {
		errorResponse = err.Error()
	}

	if errors.Is(err, ErrNoPrivilege) {
		errorResponse = err.Error()
	}

	if errors.Is(err, ErrStatusInternalError) {
		errorResponse = err.Error()
	}

	return helper.ErrorHandler(c, http.StatusNotFound, errorResponse)
}
