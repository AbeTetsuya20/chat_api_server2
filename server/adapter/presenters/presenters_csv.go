// presenters は output の interface の実装する
// 技術的要素を含んで、具体的な実装を書く。

package presenters

import (
	"server/server/entity"
	"server/server/usecase/port"
)

type CSVOutput struct {
	csvPath string
}

func (C CSVOutput) GetUsersOutputPort(users []entity.User) {
	//TODO implement me
	panic("implement me")
}

func (C CSVOutput) ErrorOutputPort(err error) {
	//TODO implement me
	panic("implement me")
}

func (C CSVOutput) LoginUserOutputPort(b bool, user *entity.User) {
	//TODO implement me
	panic("implement me")
}

func (C CSVOutput) SignUpUserOutputPort(b bool) {
	//TODO implement me
	panic("implement me")
}

func (C CSVOutput) EditProfileOutputPort(b bool) {
	//TODO implement me
	panic("implement me")
}

func NewCSVOutputPort(path string) port.ServiceOutputPort {
	return &CSVOutput{
		csvPath: path,
	}
}
