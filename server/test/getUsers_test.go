package test

//
//import (
//	"context"
//	"database/sql"
//	"encoding/json"
//	"log"
//	"net/http"
//	"net/http/httptest"
//	"server/server/driver"
//	"server/server/entity"
//	"server/server/usecase/port"
//	"testing"
//	"time"
//)
//
//type ServiceInputMock struct {
//	GatewayPort port.ServiceOutputPort
//	Repository port.ServiceRepository
//}
//
//func (s *ServiceInputMock) GetUsersInputPort(ctx context.Context) {
//	users := s.Repository.GetUsersRepository(ctx)
//	s.GatewayPort.GetUsersOutputPort(users)
//}
//
//type ServiceRepositoryMock struct {
//	now  func() time.Time
//	r    *http.Request
//	conn *sql.DB
//}
//
//func (s *ServiceRepositoryMock) GetUsersRepository(ctx context.Context) []entity.User {
//	return []entity.User{
//		{
//			ID:      "1",
//			Name:    "test",
//			Address: "",
//		},
//	}
//}
//
//type ServiceOutputMock struct {
//	w http.ResponseWriter
//}
//
//func (s *ServiceOutputMock) GetUsersOutputPort(users []entity.User) {
//	type getUser struct {
//		Name       string `json:"name"`
//		ID         string `json:"ID"`
//		ChatNumber int    `json:"chatNumber"`
//	}
//
//	res := struct {
//		User []getUser `json:"Users"`
//	}{}
//
//	var getUsers []getUser
//	for _, e := range users {
//		getUsers = append(getUsers, getUser{
//			Name:       e.Name,
//			ID:         e.ID,
//			ChatNumber: e.ChatNumber,
//		})
//	}
//
//	res.User = getUsers
//
//	if err := json.NewEncoder(s.w).Encode(&res); err != nil {
//		log.Printf("[ERROR] response encoding failed: %+v", err)
//		entity.WriteHTTPError(s.w, http.StatusInternalServerError)
//	}
//}
//
//func GetUsersTest(t *testing.T) {
//	ctx := context.Background()
//
//	outputFactory := driver.NewOutputFactory()
//	inputFactory := driver.NewInputFactory()
//	repositoryFactory := driver.NewRepositoryFactory()
//
//	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//
//	}))
//	defer ts.Close()
//
//}
//
//func (s *ServiceInputMock) LoginUserInputPort(ctx context.Context) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *ServiceInputMock) SignUpUserInputPort(ctx context.Context) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *ServiceInputMock) EditProfileInputPort(ctx context.Context) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *ServiceRepositoryMock) LoginUserRepository(ctx context.Context) (bool, *entity.User) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *ServiceRepositoryMock) SignUpUserRepository(ctx context.Context) bool {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *ServiceRepositoryMock) EditProfileRepository(ctx context.Context) bool {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *ServiceOutputMock) ErrorOutputPort(err error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *ServiceOutputMock) LoginUserOutputPort(b bool, user *entity.User) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *ServiceOutputMock) SignUpUserOutputPort(b bool) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *ServiceOutputMock) EditProfileOutputPort(b bool) {
//	//TODO implement me
//	panic("implement me")
//}
