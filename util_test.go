package pitOrgan

import (
	"golang.org/x/xerrors"
)

func nakedError(err error) error {
	if u, ok := err.(xerrors.Wrapper); ok {
		return nakedError(u.Unwrap())
	}
	return err
}
