// gateway パッケージは、DB操作に対するアダプター。
// SELECT 文などと言った、めっちゃ具体的なことを書く。

package gateway

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"server/server/entity"
	"server/server/usecase/port"
	"time"
)

type ServiceRepo struct {
	now  func() time.Time
	r    *http.Request
	conn *sql.DB
}

func NewServiceRepository(now func() time.Time, r *http.Request, conn *sql.DB) port.ServiceRepository {
	return &ServiceRepo{
		now:  now,
		r:    r,
		conn: conn,
	}
}

func ScanUsers(rows *sql.Rows) ([]entity.User, int, error) {
	users := make([]entity.User, 0)

	for rows.Next() {
		var v entity.User
		if err := rows.Scan(&v.ID, &v.Name, &v.Address, &v.Status, &v.Password, &v.ChatNumber, &v.Token, &v.CreatedAT, &v.UpdatedAt); err != nil {
			log.Printf("[ERROR] scan user: %+v", err)
			return nil, 0, err
		}
		users = append(users, v)
	}

	return users, len(users), nil
}

func (s ServiceRepo) GetUsersRepository(ctx context.Context) []entity.User {
	query := "SELECT * FROM user"
	rows, err := s.conn.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[ERROR] not found User: %+v", err)
		return nil
	}

	users, _, err := ScanUsers(rows)
	if err != nil {
		log.Printf("[ERROR] can not scan User: %+v", err)
		return nil
	}

	return users
}

func (s ServiceRepo) LoginUserRepository(ctx context.Context) (bool, *entity.User) {

	// header から情報を取得
	headerName := s.r.Header.Get("name")
	headerAddress := s.r.Header.Get("address")
	headerPassword := s.r.Header.Get("password")

	if headerName == "" || headerAddress == "" || headerPassword == "" {
		log.Printf("[ERROR] can't login: 情報が足りません。header に 名前、アドレス、パスワードがあることを確認してください。")
		return false, nil
	}

	// データベースから値を持ってくる
	query := "SELECT * FROM user WHERE name = ? AND address = ? AND password = ?"
	rows, err := s.conn.QueryContext(ctx, query, headerName, headerAddress, headerPassword)
	if err != nil {
		log.Printf("[ERROR] can't login: %+v", err)
		return false, nil
	}

	users, count, err := ScanUsers(rows)
	if count != 1 {
		log.Printf("[ERROR] can't login: 複数一致、または一致しません。")
		return false, nil
	}

	// token 生成
	var token string
	token = users[0].ID + entity.RandomWithCharset(3)

	// token を登録
	query2 := "UPDATE user SET token = ? ,updated_at = ? WHERE id = ?"
	_, err = s.conn.ExecContext(ctx, query2, token, s.now(), users[0].ID)
	if err != nil {
		log.Printf("[ERROR] can't update token: %+v", err)
		return false, nil
	}

	// token を最新のものに変更
	users[0].Token = sql.NullString{String: token}

	return true, &users[0]
}

func (s ServiceRepo) EditProfileRepository(ctx context.Context) bool {

	token := s.r.Header.Get("token")

	req := struct {
		ID             string `json:"id"`
		ProfileMessage string `json:"profile_message"`
	}{}

	if err := json.NewDecoder(s.r.Body).Decode(&req); err != nil {
		// デコードに失敗した場合はログ出力して 400 Bad Request を返す。
		log.Printf("[ERROR] request decoding failed: %+v", err)
		return false
	}

	// token が正しいか確かめる
	queryGetCountByToken := "select count(*) from user where token = ? AND id = ?"
	rows, err := s.conn.QueryContext(ctx, queryGetCountByToken, token, req.ID)
	if err != nil {
		log.Printf("[ERROR] can't get user: %+v", err)
		return false
	}

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Printf("[ERROR] can't scan count: %+v", err)
			return false
		}
	}

	if count != 1 {
		log.Printf("[ERROR] can't get user by token: 複数一致、または不正なトークンです。")
		return false
	}

	// ID に紐づく profile を更新する
	queryUpdateUserProfile := "UPDATE User_Profile SET Comment = ? ,updated_at = ? WHERE id = ?"
	_, err = s.conn.ExecContext(ctx, queryUpdateUserProfile, req.ProfileMessage, s.now(), req.ID)
	if err != nil {
		log.Printf("[ERROR] can't update Profile: %+v", err)
		return false
	}

	return true
}
