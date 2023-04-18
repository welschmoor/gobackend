package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/welschmoor/gobackend/internal/repository"
	"github.com/welschmoor/gobackend/internal/repository/dbrepo"
)

const port = 3000

type application struct {
	DSN          string //this is just the connection string
	Domain       string
	DB           repository.DatabaseRepo
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
}

func main() {
	var app application

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=newdb sslmode=disable timezone=UTC connect_timeout=5", "postgres connection string")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "RaNdOmStRiNg", "jwt-secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "example.website.comcom", "jwt-issuer")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "example.website.comcom", "jwt-audience")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "cookie-domain")
	flag.StringVar(&app.Domain, "domain", "example.website.comcom", "domain")
	flag.Parse()

	conn, err := app.connectDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 30,
		RefreshExpiry: time.Hour * 24,
		CookieDomain:  app.CookieDomain,
		CookiePath:    app.auth.CookiePath, //which app path is the cookie good for
		CookieName:    "__Host-refresh_token",
	}

	log.Printf("Starting the server on port %d ...", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes()); err != nil {
		log.Fatal(err)
	}
}
