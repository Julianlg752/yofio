// vim: sw=4 ts=4 expandtab 
package main

import (
    "fmt"
)


func Statistics(invest_list []Invest) string {
    if len(invest_list) < 1 {
        response := fmt.Sprintf("{\"Error\": \"Not enough data\"}")
        return response
    }
    total_operation := len(invest_list)
    total_fail := 0
    total_done := 0
    avg_sucesfull := 0
    avg_fail := 0
    for i  := range invest_list {
        invest := invest_list[i]
        if invest.Successful == 1 {
            avg_sucesfull += int(invest.Investment)
            total_done++
        }else {
            avg_fail += int(invest.Investment)
            total_fail++
        }
    }
    avg_sucesfull = avg_sucesfull / total_done
    avg_fail = avg_fail / total_fail
    response := fmt.Sprintf("{\"total_assignations\":%d, \"total_successful_assignations\":%d, \"total_no_successful_assingations\":%d, \"avg_successful_invest\":%d, \"avg_no_successful_invest\":%d}",
                total_operation, total_done, total_fail, avg_sucesfull, avg_fail)
    return response
}

func GetInvestments() ([]Invest, error) {
    var investment_list []Invest
    db_connection, db_err := connection()
    if db_err != nil {
        return investment_list, db_err
    }
    rows, rows_err := db_connection.Query("select * from investment")
    if rows_err != nil {
        return investment_list, rows_err
    }
    for rows.Next() {
        var investment int32
        var successful int
        scan_err := rows.Scan(&investment, &successful)
        if scan_err != nil {
            return investment_list, scan_err
        }
        investment_list = append(investment_list, Invest{ Investment: investment, Successful: successful})
    }
    return investment_list, nil
}
