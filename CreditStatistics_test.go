
// vim: sw=4 ts=4 expandtab
package main

import (
    "strings"
    "testing"
)
func TestGetInvestment(t *testing.T) {
    invest_list, invest_err := GetInvestments()
    if invest_err != nil {
        t.Error("Error in the database data: ", invest_err.Error())
    }
    t.Log("Successful ", invest_list)
}

func TestStatistics(t *testing.T) {
    var invest = []Invest{
        {
            Investment: 3000,
            Successful: 1,
        },
        {
            Investment: 4000,
            Successful: 0,
        },
        {
            Investment: 1000,
            Successful: 0,
        },
        {
            Investment: 5300,
            Successful: 1,
        },
    }
    result := Statistics(invest)
    if result == "" {
        t.Logf("Error in the statistics: %s", result)
    }
    t.Logf("Result: %s", result)
}

func TestDbStatistics(t *testing.T) {
    invest_list, invest_err := GetInvestments()
    if invest_err != nil {
        t.Logf("Error in the data from db: %s", invest_err.Error())
    }
    result := Statistics(invest_list)
    if strings.Contains(result, "Error") {
        t.Logf("Error in the statistics from db: %s", result)
    }
    t.Logf("Statistics from DB: %s ", result)
}

func TestFailStatistics(t *testing.T) {
    var invest []Invest
    result := Statistics(invest)
    if strings.Contains(result, "Error") {
        t.Logf("Error in the statistics: %s", result)
    }
}


