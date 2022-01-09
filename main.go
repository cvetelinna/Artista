package main

import (
	"Artista/common"
	"Artista/pkg/user"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
)

func main() {
	conn, err := sql.Open("pgx", os.Getenv("ARTISTA"))
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}
	defer conn.Close()

	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Printf("connected!")

	e := echo.New()
	e.Debug = true

	cfg := &common.Config{}
	err = cfg.Load()
	if err != nil {
		log.Fatalf("unable to read config: %v", err)
	}

	fmt.Println(cfg.JwtSigningSecret)

	userRepo := user.NewRepo(conn)
	userService := user.NewService(userRepo, cfg)

	user.InitRESTHandler(e, userService)

	e.Logger.Fatal(e.Start(":1323"))
}
