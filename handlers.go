// vim: sw=4 ts=4 expandtab
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func CreditAssignment(w http.ResponseWriter, r *http.Request){
    body := r.Body
    rbody, bodyErr := ioutil.ReadAll(body)
    if bodyErr != nil {
        log.Println("Body Error")
        message := "{\"Body Error\":\""+bodyErr.Error()+"\"}"
        fmt.Fprintf(w, message)
        return
    }
    var invest Invest
    unmarshalErr := json.Unmarshal(rbody, &invest)
    if unmarshalErr != nil {
        log.Println("Json Error")
        message := "{\"Cannot Parse Json\":\""+unmarshalErr.Error()+"\"}"
        fmt.Fprintf(w, message)
        return
    }
    fmt.Println(invest.Investment)
    fmt.Fprintf(w, "ping")
}
