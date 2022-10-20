package business

// to hold user data
type User struct {
	Username string
	Fullname string
	Email    string
	Password string
}

type LoginUserResponse struct {
	User      User
	SessionId string
}

// define request and response
type CreateUserRequest struct {
	Username string
	Fullname string
	Email    string
	Password string
}

type CreateUserResponse struct {
	User User
}

type LoginUserRequest struct {
	Username string
	Password string
}
