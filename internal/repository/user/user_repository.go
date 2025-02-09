package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hilmiikhsan/thrifting-app-service/constants"
	"github.com/hilmiikhsan/thrifting-app-service/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	DB     *sqlx.DB
	Logger *logrus.Logger
}

func (r *UserRepository) FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var res = new(models.User)

	err := r.DB.GetContext(ctx, res, r.DB.Rebind(queryFindUserByUsername), username)
	if err != nil {
		if err == sql.ErrNoRows {
			r.Logger.Error("repo::FindByUsername - User not found : ", err)
			return nil, errors.New(constants.ErrUsernameOrPasswordIsIncorrect)
		}

		r.Logger.Error("repo::FindByUsername - Failed to find user : ", err)

		return nil, err
	}

	return res, nil
}

func (r *UserRepository) FindUserByID(ctx context.Context, id string) (*models.User, error) {
	var res = new(models.User)

	err := r.DB.GetContext(ctx, res, r.DB.Rebind(queryFindUserByID), id)
	if err != nil {
		if err == sql.ErrNoRows {
			r.Logger.Error("repo::FindByID - User not found : ", err)
			return nil, errors.New(constants.ErrUserNotFound)
		}

		r.Logger.Error("repo::FindByID - Failed to find user : ", err)

		return nil, err
	}

	return res, nil
}
