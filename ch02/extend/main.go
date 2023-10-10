package extend

import "fmt"

type HTTPStatus int

const (
	StatusOK           HTTPStatus = 200
	StatusUnauthorized HTTPStatus = 401
)

func (s HTTPStatus) String() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusUnauthorized:
		return "Unauthorized"
	default:
		return fmt.Sprintf("HTTPStatus(%d)", s)
	}
}
