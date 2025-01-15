package usecase

type userUsecase struct{}

func NewUserUsecase() UserUsecase {
	return &userUsecase{}
}
