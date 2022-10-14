package interactor

import (
	"context"
	"server/server/entity"
	"server/server/usecase/port"
)

type DBInput struct {
	GatewayPort port.DBGatewayPort
}

func (D DBInput) GetUsersRepositoryInput(ctx context.Context) []entity.User {
	users := D.GatewayPort.GetUsersRepository(ctx)
	return users
}

func NewDBInput(outputPort port.DBGatewayPort) port.DBInputPort {
	return &DBInput{
		GatewayPort: outputPort,
	}
}
