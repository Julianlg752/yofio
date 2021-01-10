// vim: sw=4 ts=4 expandtab
package main

import (
    "testing"
    "fmt"
)

func TestFailAssign(t *testing.T) {
    invest := Invest {
        Investment: 9900,
    }
    _, _, _, credit_err := invest.Assign()
    if credit_err != nil {
        t.Errorf("Error invest: %s", credit_err)
    }
}

func TestAssign(t *testing.T) {
    invest := Invest {
        Investment: 3000,
    }
    credit_700, credit_500, credit_300, credit_err := invest.Assign()
    if credit_err != nil {
        t.Errorf("Error invest: %s", credit_err)
    }
    response := fmt.Sprintf("{\"credit_type_300\": %d, \"credit_type_500\": %d, \"credit_type_700\": %d}", credit_300, credit_500, credit_700)
    t.Log(response)
}

func TestSaveInvestmentFail(t *testing.T) {
    query := ""
    save_err := SaveInvestment(query)
    if save_err != nil {
        t.Log("Error Saving Empty Query: ", save_err)
    }
}

func TestSaveInvestment(t *testing.T) {
    query := "INSERT INTO investment (investment, successfull) values(6700, 1)"
    save_err := SaveInvestment(query)
    if save_err != nil {
        t.Error("Error Saving Empty Query: ", save_err)
    }
    t.Log("Investment Saved")
}

