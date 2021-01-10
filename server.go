// vim: sw=4 ts=4 expandtab
package main

import (
    "fmt"
    "net/http"
    "log"
    "os"

    "github.com/gorilla/mux"
)
func main(){
    r := mux.NewRouter()
    r.HandleFunc("/credit-assignment", CreditAssignment)
    r.HandleFunc("/statistics", InvestmentStatistics).Methods("POST")
    PORT := os.Getenv("PORT")
    log.Println("Serves is alive in port:"+PORT)
    if err := http.ListenAndServe(":"+PORT, r); err != nil {
        fmt.Println("Server Error", err)
        return
    }
}
