package db

import "refty-backend/internal/domain/repository"

type userRepo struct{}

func NewUserRepo() repository.UserRepo {
	return &userRepo{}
}
