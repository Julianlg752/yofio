// vim: sw=4 ts=4 expandtab
package main

import (
    "errors"
    "fmt"
    "log"
    "math"
    "sort"
)
//Invest struct represent the amount of the investment also is used to get the Data from the Database
type Invest struct {
    Investment int32
    Successful int
}

var OK = 1
var ERROR = 0
type CreditAssigner interface {
    Assign(int32) (int32, int32, int32, error)
}

//Assign
// This function use the investmnent and assing this to a diferent credits
// that are located in 700, 500, 300 
// Iterate hover the all possible combinations to determine which one is the correct and check that the all investment 
// was used. Then return the size of the credits if was possible or return an error if couldn't proccess the investment.
// Also save The result in the database to generate the statistics
func (i Invest) Assign() (int32, int32, int32, error) {
    creditos := []int{700, 500, 300}
    inversion := int(i.Investment)
    cantidad_credito := 0
    total_creditos := make(map[int]int)
    rotation := 0
    temp := creditos
    cont := 0
    tries := 0
    for {
        if inversion == 0 {
            break
        }else {
            cantidad_credito++
            for i := range creditos {
                credito := creditos[i]
                inversion = inversion - credito
                total_creditos[credito] = cantidad_credito
                if inversion == 0 {
                    break
                }
                if math.Signbit(float64(inversion)) {
                    inversion = inversion + credito
                    if tries == len(creditos) {
                        total_creditos[-1] = -1
                        inversion = 0
                        break
                    }
                    tries++
                }
            }
            var newarray []int
            for i := 0; i <= rotation; i++ {
                newarray = temp[1:len(creditos)]
                newarray = append(newarray, temp[0])
                temp = newarray
            }
            if cont == 3 {
                sort.Ints(temp)
                rotation++
            }
            creditos = temp
            cont++
        }
    }
    for k  := range total_creditos {
        if k == -1 {
            query := fmt.Sprintf("INSERT INTO investment values(%d,%d)",i.Investment, ERROR)
            SaveInvestment(query)
            return 0, 0, 0, errors.New("Couldn't Proccess the Investment")
        }
    }
    query := fmt.Sprintf("INSERT INTO investment values(%d,%d)",i.Investment, OK)
    SaveInvestment(query)
    return int32(total_creditos[700]), int32(total_creditos[500]), int32(total_creditos[300]), nil
}

//SaveInvestment
// This function recives a query and sent it to the database to save the result of the investment
func SaveInvestment(query string) error {
    db_connection, db_err := Connection()
    if db_err != nil {
        log.Println(db_err)
        return db_err
    }
    _, result_err := db_connection.Exec(query)
    if result_err != nil {
        log.Println("result_err: ", result_err)
        return result_err
    }
    db_connection.Close()
    return nil
}
