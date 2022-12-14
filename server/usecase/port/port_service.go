package port

import (
	"context"
	"server/server/entity"
)

// port には interface の定義だけする
// 実装は interactor に書く。
// ここには技術的要素を含む、具体的な実装は書いてはいけない

// 1. 情報を受け取る
type ServiceInputPort interface {
	GetUsersInputPort(context.Context)
	LoginUserInputPort(context.Context)
	SignUpUserInputPort(ctx context.Context)
	EditProfileInputPort(context.Context)
}

type ServiceRepository interface {
	GetUsersRepository(context.Context) []entity.User
	LoginUserRepository(context.Context) (bool, *entity.User)
	SignUpUserRepository(ctx context.Context) bool
	EditProfileRepository(ctx context.Context) bool
}

// 3. 情報を出力する
type ServiceOutputPort interface {
	GetUsersOutputPort([]entity.User)
	ErrorOutputPort(error)
	LoginUserOutputPort(bool, *entity.User)
	SignUpUserOutputPort(bool)
	EditProfileOutputPort(bool)
}
