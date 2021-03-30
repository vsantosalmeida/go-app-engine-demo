package person

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

func NewErrPersonNotFound() *ErrPersonNotFound {
	return &ErrPersonNotFound{
		msg: "person not found",
	}
}

func NewErrValidatePerson() *ErrValidatePerson {
	return &ErrValidatePerson{
		msg: "person with age less than 18 must have a valid parentkey",
	}
}

func NewErrDeletePerson() *ErrDeletePerson {
	return &ErrDeletePerson{
		msg: "couldn't delete a person with an active parentkey",
	}
}
