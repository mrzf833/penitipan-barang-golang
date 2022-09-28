package exception

type AuthorizationError struct {
}

func NewAuthorizationError() AuthorizationError {
	return AuthorizationError{}
}
