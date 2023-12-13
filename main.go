package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main () {




	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if(portString == ""){
		log.Fatal("port is not found in the env file")
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	router.Mount("/v1",v1Router)


	fmt.Println("Running At Port:", portString)
	 srv:= &http.Server{
		Handler: router,
		Addr: ":" + portString,
	 }

	 err := srv.ListenAndServe()
	 if err != nil {
		log.Fatal(err)
	 }


}