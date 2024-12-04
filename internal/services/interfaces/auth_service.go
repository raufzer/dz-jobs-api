package interfaces

import "dz-jobs-api/data/request"

type AuthService interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUsersRequest) error
}
