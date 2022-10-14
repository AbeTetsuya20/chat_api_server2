package gateway

import (
	"context"
	"database/sql"
	"log"
	"server/server/entity"
	"server/server/usecase/port"
)

type DBGatewayPort struct {
	conn *sql.DB
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

func (s DBGatewayPort) GetUsersRepository(ctx context.Context) []entity.User {
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

func NewDBGatewayPort(conn *sql.DB) port.DBGatewayPort {
	return &DBGatewayPort{
		conn: conn,
	}
}
