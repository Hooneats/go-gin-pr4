package common

type Error struct {
	Code int
	Err  error
}

func NewError(e error, c int) Error {
	return Error{
		Code: c,
		Err:  e,
	}
}
