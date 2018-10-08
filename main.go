package main

import (
	"fmt"
	"github.com/zxxz-ru/go-image-server/server"
	"github.com/zxxz-ru/go-image-server/utils"
	"log"
	"net/http"
)

func main() {
	fmt.Println(utils.Prt())
	fmt.Println("Server started. Listening port 8081.")
	http.HandleFunc("/images/", server.ImagesHandler)
	http.HandleFunc("/", server.HomeHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
