// vim: sw=4 ts=4 expandtab
package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

type Invest struct {
    Investment int32
}

type CreditAssigner interface {
    Assign(int32) (int32, int32, int32, error)
}

func (i Invest) Assign() (int32, int32, int32, error) {
	creditos := []int{700, 500, 300}
	inversion := int(i.Investment)
	cantidad_credito := 0
	total_creditos := make(map[int]int)
	rotation := 0
	temp := creditos
	cont := 0
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
					inversion = 0
					total_creditos[-1] = -1
					break
				}
			}
			var newarray []int
			for i := 0; i <= rotation; i++ {
				newarray = temp[1:len(creditos)]
				newarray = append(newarray, temp[0])
				temp = newarray
			}
			if cont == 3 {
				fmt.Println()
				sort.Ints(temp)
				rotation++
			}
			creditos = temp
			cont++

		}
	}
    for k  := range total_creditos {
        if k == -1 {
            return 0, 0, 0, errors.New("Couldn't Proccess the Investment")
        }
    }
    return int32(total_creditos[700]), int32(total_creditos[500]), int32(total_creditos[300]), nil
}

