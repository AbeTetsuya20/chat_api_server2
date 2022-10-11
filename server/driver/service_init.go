package driver

import (
	"database/sql"
	"server/server/adapter/controller"
	"server/server/adapter/gateway"
	"server/server/adapter/presenters"
	"server/server/usecase/interactor"
)

func InitDriver(con *sql.DB) (Service, error) {
	outputFactory := NewOutputFactory()
	inputFactory := NewInputFactory()
	repositoryFactory := NewRepositoryFactory()
	user := controller.NewServiceController(inputFactory, repositoryFactory, outputFactory, con)
	driversUser := NewServiceDriver(user)
	return driversUser, nil
}

func NewOutputFactory() controller.OutputFactory {
	return presenters.NewServiceOutputPort
	//return presenters.NewCSVOutputPort
}

func NewInputFactory() controller.InputFactory {
	return interactor.NewServiceInputPort
}

func NewRepositoryFactory() controller.Repository {
	return gateway.NewServiceRepository
}
