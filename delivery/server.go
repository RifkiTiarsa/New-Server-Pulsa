package delivery

import (
	"database/sql"
	"fmt"
	"server-pulsa/config"
	"server-pulsa/delivery/controller"
	"server-pulsa/repository"
	"server-pulsa/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	memberUc usecase.MemberUsecase
	engine   *gin.Engine
	host     string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)
	controller.NewMemberController(s.memberUc, rg).Routes()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running on host %s, becauce error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()
	dsn := fmt.Sprintf("%s:%s@/%s?parseTime=true", cfg.Username, cfg.Password, cfg.DBName)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic(fmt.Errorf("failed to connect to database, error %v", err.Error()))
	}

	memberRepo := repository.NewMemberRepository(db)
	memberUc := usecase.NewMemberUsecase(memberRepo)
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		memberUc: memberUc,
		engine:   gin.Default(),
		host:     host,
	}
}
