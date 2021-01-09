// vim: sw=4 ts=4 expandtab
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)
func main(){
    r := mux.NewRouter()
    PORT := os.Getenv("PORT")
    if err := http.ListenAndServe(":"+PORT, r); err != nil {
        fmt.Println("Error Iniciando Servidor", err)
        return
    }
}
