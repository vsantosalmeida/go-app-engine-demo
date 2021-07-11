package person

import (
	"cloud.google.com/go/datastore"
)

func handleError(err error) error {
	switch err {
	case datastore.ErrNoSuchEntity:
		return NewErrPersonNotFound(err.Error())
	default:
		if err.Error() == "" {

		}
		return err
	}
}
