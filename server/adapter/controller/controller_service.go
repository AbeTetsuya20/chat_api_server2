// entity -> port -> interactor -> presenters -> gateway -> controller
// input port に情報を渡す。
// 技術的な実装を記述する

package controller

import (
	"context"
	"database/sql"
	"net/http"
	"server/server/usecase/port"
	"time"
)

type ServiceController struct {
	InputFactory     func(o port.ServiceOutputPort, u port.ServiceRepository, inputPort port.DBInputPort, outputPort port.DBGatewayPort) port.ServiceInputPort
	RepoFactory      func(now func() time.Time, r *http.Request, c *sql.DB) port.ServiceRepository
	DBInputFactory   func(o port.DBGatewayPort) port.DBInputPort
	DBGatewayFactory func(c *sql.DB) port.DBGatewayPort
	OutputFactory    func(w http.ResponseWriter) port.ServiceOutputPort
	//OutputFactory func() port.ServiceOutputPort

	Conn *sql.DB
}

type InputFactory func(o port.ServiceOutputPort, u port.ServiceRepository, inputPort port.DBInputPort, outputPort port.DBGatewayPort) port.ServiceInputPort
type Repository func(now func() time.Time, r *http.Request, c *sql.DB) port.ServiceRepository
type DBInputFactory func(o port.DBGatewayPort) port.DBInputPort
type DBGatewayFactory func(c *sql.DB) port.DBGatewayPort

type OutputFactory func(w http.ResponseWriter) port.ServiceOutputPort

//type OutputFactory func() port.ServiceOutputPort

func NewServiceController(inputFactory InputFactory, dbInputFactory DBInputFactory, dbOutputFactory DBGatewayFactory, repository Repository, outputFactory OutputFactory, conn *sql.DB) *ServiceController {
	return &ServiceController{
		InputFactory:     inputFactory,
		RepoFactory:      repository,
		OutputFactory:    outputFactory,
		DBInputFactory:   dbInputFactory,
		DBGatewayFactory: dbOutputFactory,
		Conn:             conn,
	}
}

func (s *ServiceController) NewInputPort(w http.ResponseWriter, r *http.Request) port.ServiceInputPort {
	outputPort := s.OutputFactory(w)
	repository := s.RepoFactory(time.Now, r, s.Conn)
	dbOutput := s.DBGatewayFactory(s.Conn)
	dbInput := s.DBInputFactory(dbOutput)
	return s.InputFactory(outputPort, repository, dbInput, dbOutput)
}

//func (s *ServiceController) NewInputPort(w http.ResponseWriter, r *http.Request) port.ServiceInputPort {
//	outputPort := s.OutputFactory(w)
//	repository := s.RepoFactory(time.Now, r, s.Conn)
//	return s.InputFactory(outputPort, repository)
//}

func (s *ServiceController) GetUsers(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		s.NewInputPort(w, r).GetUsersInputPort(ctx)
	}

	return handleFunc
}

func (s *ServiceController) LoginUser(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		s.NewInputPort(w, r).LoginUserInputPort(ctx)
	}

	return handleFunc
}

func (s *ServiceController) EditProfile(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		s.NewInputPort(w, r).EditProfileInputPort(ctx)
	}

	return handleFunc
}

func (s *ServiceController) SignUpUser(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		s.NewInputPort(w, r).SignUpUserInputPort(ctx)
	}

	return handleFunc
}
