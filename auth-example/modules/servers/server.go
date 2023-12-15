package servers

import (
	"auth-example/configs"
	"auth-example/pkg/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	App *fiber.App
	Cfg *configs.Configs
	Db  *sqlx.DB
}

func NewServer(cfg *configs.Configs, db *sqlx.DB) *Server {
	return &Server{
		App: fiber.New(),
		Cfg: cfg,
		Db:  db,
	}
}

func (s *Server) Start() {
	if err := s.MapHandlers(); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	fibrtConnURL, err := utils.ConnectionUrlBuilder("fiber", s.Cfg)
	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	host := s.Cfg.App.Host
	port := s.Cfg.App.Post
	log.Printf("server has been started on %s:%s ⚡️", host, port)

	if err := s.App.Listen(fibrtConnURL); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
}
