package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AJ-Brown-InTech/libre-api/config"
	"github.com/AJ-Brown-InTech/libre-api/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
)

type Server struct{
	echo *echo.Echo
	logs utils.Logger
	db *sql.DB
	cfg *config.Config
}

const (
	readtimeout = time.Second * 8
	writetimeout = readtimeout
	maxheaderbytes = 1 << 20
)

//new server constructor
func NewServer(c *config.Config, d *sql.DB, l utils.Logger) *Server{
	return &Server{echo: echo.New(), cfg: c, db: d, logs: l  }
}

func (s *Server) Run()error{
	
	server := &http.Server{
		Addr: s.cfg.Server.Port,
		ReadTimeout: readtimeout,
		WriteTimeout: writetimeout,
		MaxHeaderBytes: maxheaderbytes,
	}

	go func(){
	s.logs.Infof("Server listening on PORT: %s",s.cfg.Server.Port)	
		if err := s.echo.StartServer(server); err != nil{
			s.logs.Errorf("Error starting server: ", err)
		}
	}()

	go func(){
		s.logs.Infof("Debugging server starting on PORT: %s", s.cfg.Server.PprofPort)
		if err := http.ListenAndServe(s.cfg.Server.PprofPort, http.DefaultServeMux); err != nil {
			s.logs.Errorf("Error statring Debugging server: ", err)
		}
	}()
	

	quit := make(chan os.Signal,1)
	signal.Notify(quit,os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx ,shutdown := context.WithTimeout(context.Background(), 5 * time.Second)
	defer shutdown()

	s.logs.Infof("Server Exited Properly %v",ctx)
		return s.echo.Server.Shutdown(ctx)
	
}


func Serve(){
	e := echo.New()
	e.GET("/", func(c echo.Context)error{
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
	fmt.Println(("server here."))
}