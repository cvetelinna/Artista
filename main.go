package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	conn, err := sql.Open("pgx", os.Getenv("ARTISTA"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close()

	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected!")

	//userRepo := users.NewRepo(conn)
	//
	//e := echo.New()

}
