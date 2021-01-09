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
    r.HandleFunc("/credit-assignment", CreditAssignment)
    PORT := os.Getenv("PORT")
    if err := http.ListenAndServe(":"+PORT, r); err != nil {
        fmt.Println("Server Error", err)
        return
    }
}
