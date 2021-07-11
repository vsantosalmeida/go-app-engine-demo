package person

import "fmt"

type ErrPersonNotFound struct {
	msg string
}

func (e *ErrPersonNotFound) Error() string { return e.msg }

type ErrValidatePerson struct {
	msg string
}

func (e *ErrValidatePerson) Error() string { return e.msg }

type ErrDeletePerson struct {
	msg string
}

func (e *ErrDeletePerson) Error() string { return e.msg }

func NewErrPersonNotFound(reason string) *ErrPersonNotFound {
	return &ErrPersonNotFound{
		msg: fmt.Sprintf("person not found, reason=%s", reason),
	}
}

func NewErrValidatePerson(reason string) *ErrValidatePerson {
	return &ErrValidatePerson{
		msg: fmt.Sprintf("failed to validate person, reason=%s", reason),
	}
}

func NewErrDeletePerson(reason string) *ErrDeletePerson {
	return &ErrDeletePerson{
		msg: fmt.Sprintf("couldn't delete person, reason=%s", reason),
	}
}
