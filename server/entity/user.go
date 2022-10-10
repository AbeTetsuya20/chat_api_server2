package entity

import (
	"crypto/rand"
	"database/sql"
	"math/big"
	"time"
)

type User struct {
	ID         string
	Name       string
	Address    string
	Status     string
	Password   string
	ChatNumber int
	Token      sql.NullString
	CreatedAT  time.Time
	UpdatedAt  time.Time
}

type Users []User

func RandomWithCharset(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)

	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}

		b[i] = charset[n.Int64()]
	}

	return string(b)
}
