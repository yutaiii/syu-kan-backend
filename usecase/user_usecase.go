package usecase

import (
	"context"

	"github.com/yutaiii/syu-kan-backend/domain/model"
	"github.com/yutaiii/syu-kan-backend/repository"
	"github.com/yutaiii/syu-kan-backend/tool/util"
)

type UserUsecase struct {
	ctx        context.Context
	repository *repository.UserRepository
}

func NewUserUsecase(ctx context.Context) *UserUsecase {
	return &UserUsecase{
		ctx:        ctx,
		repository: repository.NewUserRepository(ctx),
	}
}

func (u *UserUsecase) CreateUser(model *model.InputUser) error {
	db := util.GetConn()
	return u.repository.CreateUser(db, model)
}

func (u *UserUsecase) FindUserByFirebaseUID(model *model.InputUser) (*model.OutputUser, error) {
	db := util.GetConn()
	return u.repository.FindUserByFirebaseUID(db, model)
}
