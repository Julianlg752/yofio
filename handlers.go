// vim: sw=4 ts=4 expandtab
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func InvestmentStatistics(w http.ResponseWriter, r *http.Request){
    investment_list, investment_err := GetInvestments()
    if investment_err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, investment_err.Error())
        return
    }
    statistics := Statistics(investment_list)
    fmt.Fprintf(w, statistics)
}

func CreditAssignment(w http.ResponseWriter, r *http.Request){
    body := r.Body
    rbody, bodyErr := ioutil.ReadAll(body)
    if bodyErr != nil {
        log.Println("Body Error")
        message := "{\"Body Error\":\""+bodyErr.Error()+"\"}"
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, message)
        return
    }
    var invest Invest
    unmarshalErr := json.Unmarshal(rbody, &invest)
    if unmarshalErr != nil {
        log.Println("Json Error: ", unmarshalErr)
        message := "{\"Cannot Parse Json\":\""+unmarshalErr.Error()+"\"}"
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, message)
        return
    }
    credit_700, credit_500, credit_300, credit_err := invest.Assign()
    if credit_err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    response := fmt.Sprintf("{\"credit_type_300\": %d, \"credit_type_500\": %d, \"credit_type_700\": %d}", credit_300, credit_500, credit_700)
    fmt.Fprintf(w, response)
}
