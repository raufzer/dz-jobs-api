package interfaces

import "dz-jobs-api/internal/dto/request"

type AuthService interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUsersRequest) error
}
