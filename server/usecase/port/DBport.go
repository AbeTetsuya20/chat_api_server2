package port

import (
	"context"
	"server/server/entity"
)

type DBInputPort interface {
	GetUsersRepositoryInput(context.Context) []entity.User
}

type DBGatewayPort interface {
	GetUsersRepository(context.Context) []entity.User
}
