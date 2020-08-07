package gresp

import (
	"fmt"
	"github.com/DowneyL/now/packages/locales"
)

type Code uint32

const (
	OK     Code = 0
	Failed Code = 10000 + iota
	Canceled
	Unknown
	InvalidArgument
	DeadlineExceeded
	NotFound
	AlreadyExists
	PermissionDenied
	ResourceExhausted
	Aborted
	OutOfRange
	Internal
	Unavailable
	DataLoss
	Unauthenticated
)

var codeToStr = map[Code]string{
	OK:                "ok",
	Canceled:          "canceled",
	Failed:            "failed",
	Unknown:           "unknown",
	InvalidArgument:   "invalid_argument",
	DeadlineExceeded:  "deadline_exceeded",
	NotFound:          "not_found",
	AlreadyExists:     "already_exists",
	PermissionDenied:  "permission_denied",
	ResourceExhausted: "resource_exhausted",
	Aborted:           "aborted",
	OutOfRange:        "out_of_range",
	Internal:          "internal",
	Unavailable:       "unavailable",
	DataLoss:          "data_loss",
	Unauthenticated:   "unauthenticated",
}

func (code Code) String() string {
	messageId := fmt.Sprintf("response.%s", codeToStr[code])
	return locales.Translator.MustTrans(messageId)
}
