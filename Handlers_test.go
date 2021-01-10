// vim: sw=4 ts=4 expandtab
package main

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/gorilla/mux"
)

func makeRequest(investment string) *http.Request {
    request, _ := http.NewRequest(http.MethodGet, "/credit-assignment", nil)
    return request
}

var URL = "http://localhost:8000"
func TestCreditAssignment(t *testing.T) {
    r := mux.NewRouter()
    r.HandleFunc("/credit-assignment", CreditAssignment)
    ts := httptest.NewServer(r)
    defer ts.Close()
    rq, err := http.NewRequest("GET", URL+"/credit-assignment", strings.NewReader("{\"investment\":3000}"))
    if err != nil {
        t.Logf("Internal Server Error: %s", err.Error())
    }
    client := http.Client{}
    resp, err := client.Do(rq)
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Internal Server Error: %d", resp.StatusCode)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Error("Response Error: ", err.Error())
    }
    t.Log(string(body))
    t.Log(resp.StatusCode)
}

func TestStatisticsHttp(t *testing.T) {
    r := mux.NewRouter()
    r.HandleFunc("/statistics", InvestmentStatistics)
    ts := httptest.NewServer(r)
    defer ts.Close()
    rq, err := http.NewRequest("POST", URL+"/statistics", nil)
    if err != nil {
        t.Logf("Internal Server Error: %s", err.Error())
    }
    client := http.Client{}
    resp, err := client.Do(rq)
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Internal Server Error: %d", resp.StatusCode)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Error("Response Error: ", err.Error())
    }
    t.Log(string(body))
    t.Log(resp.StatusCode)

}
