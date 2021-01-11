// vim: sw=4 ts=4 expandtab
package main

import (
    "fmt"
    "net/http"
    "log"
    "os"

    "github.com/gorilla/mux"
)

// main()
// Declaring the paths of the endpoint
// /credit-assignment use it to sent a investment and get the possibles values of the credits
// usage: 
//      curl --location --request GET 'localhost:8000/credit-assignment' --data-raw '{"investment": 2500}'
// return: 
//      if the investment could be processed return the JSON with the values of the each credit 
//      or return the status code BADREQUEST if the investment couldn't processed

// /statistics use this endpoint to get a JSON with the all information of the investment in the past
// usage: 
//      curl --location --request POST 'localhost:8000/statistics'
// return:
//      Return a Json with the resume of the operations like: 
//      total_operation, successful_operations, failed_operations, avg_successful_investment, avg_no_successful_investment
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
