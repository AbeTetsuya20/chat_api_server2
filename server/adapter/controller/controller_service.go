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
	InputFactory  func(o port.ServiceOutputPort, u port.ServiceRepository) port.ServiceInputPort
	RepoFactory   func(now func() time.Time, r *http.Request, c *sql.DB) port.ServiceRepository
	OutputFactory func(w http.ResponseWriter) port.ServiceOutputPort
	Conn          *sql.DB
}

type InputFactory func(o port.ServiceOutputPort, u port.ServiceRepository) port.ServiceInputPort
type Repository func(now func() time.Time, r *http.Request, c *sql.DB) port.ServiceRepository
type OutputFactory func(w http.ResponseWriter) port.ServiceOutputPort

func NewServiceController(inputFactory InputFactory, repository Repository, outputFactory OutputFactory, conn *sql.DB) *ServiceController {
	return &ServiceController{
		InputFactory:  inputFactory,
		RepoFactory:   repository,
		OutputFactory: outputFactory,
		Conn:          conn,
	}
}

func (s *ServiceController) newInputPort(w http.ResponseWriter, r *http.Request) port.ServiceInputPort {
	outputPort := s.OutputFactory(w)
	repository := s.RepoFactory(time.Now, r, s.Conn)
	return s.InputFactory(outputPort, repository)
}

func (s *ServiceController) GetUsers(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		s.newInputPort(w, r).GetUsersInputPort(ctx)
	}

	return handleFunc
}

func (s *ServiceController) LoginUser(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		s.newInputPort(w, r).LoginUserInputPort(ctx)
	}

	return handleFunc
}

func (s *ServiceController) EditProfile(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		s.newInputPort(w, r).EditProfileInputPort(ctx)
	}

	return handleFunc
}
