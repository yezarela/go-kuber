package utils

import (
	"fmt"

	"github.com/pkg/errors"
)

type stacktracer interface {
	StackTrace() errors.StackTrace
}

type causer interface {
	Cause() error
}

func PrintStackTrace(err error) {

	var errStack errors.StackTrace

	for err != nil {
		// Find the earliest error.StackTrace
		if t, ok := err.(stacktracer); ok {
			errStack = t.StackTrace()
		}
		if c, ok := err.(causer); ok {
			err = c.Cause()
		} else {
			break
		}
	}
	if errStack != nil {
		fmt.Println(err)
		fmt.Printf("%+v\n", errStack)
	} else {
		fmt.Printf("%+v\n", errors.WithStack(err))
	}
}
