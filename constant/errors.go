package constant

const (
	ErrDuplicateEmailMessage = `ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)`
)

var (
	ErrSystemConvertTimeDuration = "Cannot convert time duration from string to int fail!"

	ErrHashCode       = "system cannot hash this"
	ErrCreateUserFail = "Create user fail!"
	CreateUserSuccess = "Create user success!"

	ErrorEmailExist    = "This account already exists!"
	InvalidRequestBody = "Invalid Request Body!"
	GetUserSuccess     = "Get user successfully"
	GetUsersSuccess    = "Get users successfully"
	DeleteUserSuccess  = "Delete user successfully"
	UpdateUserSuccess  = "Update user successfully"
)
