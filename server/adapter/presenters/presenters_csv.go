// presenters は output の interface の実装する
// 技術的要素を含んで、具体的な実装を書く。

package presenters

import (
	"fmt"
	"server/server/entity"
	"server/server/usecase/port"
)

type CSVOutput struct {
}

func (C CSVOutput) GetUsersOutputPort(users []entity.User) {
	fmt.Println("csv 出力をする")
}

func (C CSVOutput) ErrorOutputPort(err error) {
	fmt.Println("csv 出力をする")
}

func (C CSVOutput) LoginUserOutputPort(b bool, user *entity.User) {
	fmt.Println("csv 出力をする")
}

func (C CSVOutput) SignUpUserOutputPort(b bool) {
	fmt.Println("csv 出力をする")
}

func (C CSVOutput) EditProfileOutputPort(b bool) {
	fmt.Println("csv 出力をする")
}

func NewCSVOutputPort() port.ServiceOutputPort {
	return &CSVOutput{}
}
