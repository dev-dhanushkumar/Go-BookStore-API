package main

import (
	"fmt"
	"log"
	"net/http"
	"sam0307/go-bookstore/pkg/router"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	router.RegisterBookStoreRouter(r)
	http.Handle("/", r)
	fmt.Println("Library are open....")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}
