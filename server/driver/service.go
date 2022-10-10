// entity -> port -> interactor -> presenters -> gateway -> controller -> driver
// controller を呼び出す。

// driver 層に書くこと
// 1. DB の初期化
// 2. API server の起動
// 3. usecase の作成

package driver

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"net"
	"net/http"
	"os"
	"server/server/adapter/controller"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Service interface {
	ChatService(ctx context.Context)
}

type ServiceDriver struct {
	// now 現在時刻を取得するための関数
	Now        func() time.Time
	Controller *controller.ServiceController
}

func NewServiceDriver(service *controller.ServiceController) Service {
	return &ServiceDriver{
		Now:        time.Now,
		Controller: service,
	}
}

func (s *ServiceDriver) ChatService(ctx context.Context) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/debug", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"message": "hello",
		}
		render.JSON(w, r, data)
	})

	r.Route("/api", func(r chi.Router) {

		// GET: /api/users
		r.Get("/users", s.Controller.GetUsers(ctx))

		r.Route("/login", func(r chi.Router) {
			// GET: /api/login/user
			r.Get("/user", s.Controller.LoginUser(ctx))
		})

		// POST: /api/user/profile
		r.Post("/user/profile", s.Controller.EditProfile(ctx))
	})

	addr := os.Getenv("Addr")
	if addr == "" {
		addr = ":1001"
	}

	log.Printf("listen: %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("!! %+v", err)
	}
}

func InitDatabase(host string, port uint16, dbname, username, password string) (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.Addr = net.JoinHostPort(host, strconv.Itoa(int(port)))
	cfg.DBName = dbname
	cfg.User = username
	cfg.Passwd = password
	cfg.ParseTime = true

	connector, err := mysql.NewConnector(cfg)
	if err != nil {
		return nil, fmt.Errorf("new connector: %w", err)
	}

	return sql.OpenDB(connector), nil
}
