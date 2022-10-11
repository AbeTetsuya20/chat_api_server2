// presenters は output の interface の実装する
// 技術的要素を含んで、具体的な実装を書く。

package presenters

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"server/server/entity"
	"server/server/usecase/port"
)

type ServiceOutput struct {
	w http.ResponseWriter
}

func NewServiceOutputPort(w http.ResponseWriter) port.ServiceOutputPort {
	return &ServiceOutput{
		w: w,
	}
}

func (s *ServiceOutput) ErrorOutputPort(err error) {
	entity.WriteErrorResponse(s.w, 500, entity.NewError(err))
}

func (s *ServiceOutput) GetUsersOutputPort(users []entity.User) {

	type getUser struct {
		Name       string         `json:"name"`
		ID         string         `json:"ID"`
		Token      sql.NullString `json:"Token"`
		ChatNumber int            `json:"chatNumber"`
	}

	res := struct {
		User []getUser `json:"Users"`
	}{}

	var getUsers []getUser
	for _, e := range users {
		getUsers = append(getUsers, getUser{
			Name:       e.Name,
			ID:         e.ID,
			Token:      e.Token,
			ChatNumber: e.ChatNumber,
		})
	}

	res.User = getUsers

	if err := json.NewEncoder(s.w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		entity.WriteHTTPError(s.w, http.StatusInternalServerError)
	}
}

func (s *ServiceOutput) LoginUserOutputPort(success bool, user *entity.User) {

	res := struct {
		Success bool           `json:"success"`
		ID      string         `json:"id"`
		Token   sql.NullString `json:"token"`
	}{
		Success: false,
		ID:      "",
		Token:   sql.NullString{String: "", Valid: false},
	}
	if user != nil {
		res.Success = true
		res.ID = user.ID
		res.Token = user.Token
	}

	if err := json.NewEncoder(s.w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		entity.WriteHTTPError(s.w, http.StatusInternalServerError)
	}
}

func (s *ServiceOutput) EditProfileOutputPort(success bool) {
	res := struct {
		Success bool `json:"success"`
	}{
		Success: success,
	}

	if err := json.NewEncoder(s.w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		entity.WriteHTTPError(s.w, http.StatusInternalServerError)
	}
}

func (s *ServiceOutput) SignUpUserOutputPort(success bool) {
	res := struct {
		Success bool `json:"success"`
	}{
		Success: success,
	}

	if err := json.NewEncoder(s.w).Encode(&res); err != nil {
		log.Printf("[ERROR] response encoding failed: %+v", err)
		entity.WriteHTTPError(s.w, http.StatusInternalServerError)
	}
}
