package gateways

import (
	"github.com/gbroveri/users/domains"
)

type UserGateway interface {
	Save(user domains.User) (domains.User, error)
	Get(id string) (domains.User, error)
}
