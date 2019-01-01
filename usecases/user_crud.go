package usecases

import (
	"context"

	"github.com/gbroveri/users/domains"
	"github.com/gbroveri/users/gateways"
)

type userCrud struct {
	gateway gateways.UserGateway
}

type UserUsecases interface {
	Create(ctx context.Context, user domains.User) (domains.User, error)
	Get(ctx context.Context, id string) (domains.User, error)
}

func NewUserCrud(g gateways.UserGateway) UserUsecases {
	return &userCrud{gateway: g}
}

func (crud *userCrud) Create(ctx context.Context, user domains.User) (domains.User, error) {
	user.Identifiers = []string{user.Email, user.Username}
	if saved, err := crud.gateway.Save(user); err != nil {
		return domains.User{}, err
	} else {
		return saved, nil
	}
}

func (crud *userCrud) Get(ctx context.Context, id string) (domains.User, error) {
	if found, err := crud.gateway.Get(id); err != nil {
		return domains.User{}, err
	} else {
		return found, nil
	}
}
