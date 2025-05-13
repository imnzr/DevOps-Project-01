package web

type UserUpdateRequest struct {
	Id       int
	Username string
	Email    string
	Password string
}
