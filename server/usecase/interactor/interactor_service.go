// interator は input の interface の実装をする
// port で実装した output や repo の手順を実装する
// port と併せて、ここでも技術的な実装は書かない。

package interactor

import (
	"context"
	"server/server/usecase/port"
)

type ServiceInput struct {
	OutputPort port.ServiceOutputPort
	Repository port.ServiceRepository
}

func NewServiceInputPort(outputPort port.ServiceOutputPort, userRepository port.ServiceRepository) port.ServiceInputPort {
	return &ServiceInput{
		OutputPort: outputPort,
		Repository: userRepository,
	}
}

// User を取得する抽象度が高い関数
// repo から User を取得 -> output port に渡す
func (s *ServiceInput) GetUsersInputPort(ctx context.Context) {
	users := s.Repository.GetUsersRepository(ctx)
	s.OutputPort.GetUsersOutputPort(users)
}

// User Login をする関数
// 技術的なことは書かない
func (s *ServiceInput) LoginUserInputPort(ctx context.Context) {
	success, user := s.Repository.LoginUserRepository(ctx)
	s.OutputPort.LoginUserOutputPort(success, user)
}

func (s *ServiceInput) EditProfileInputPort(ctx context.Context) {
	success := s.Repository.EditProfileRepository(ctx)
	s.OutputPort.EditProfileOutputPort(success)
}
